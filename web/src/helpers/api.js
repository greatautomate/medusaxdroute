import { getUserIdFromLocalStorage, showError, formatMessageForAPI, isValidMessage } from './utils';
import axios from 'axios';
import { MESSAGE_ROLES } from '../constants/playground.constants';

export let API = axios.create({
  baseURL: import.meta.env.VITE_REACT_APP_SERVER_URL
    ? import.meta.env.VITE_REACT_APP_SERVER_URL
    : '',
  headers: {
    'MedusaXD-API-User': getUserIdFromLocalStorage(),
    'Cache-Control': 'no-store',
  },
});

function patchAPIInstance(instance) {
  const originalGet = instance.get.bind(instance);
  const inFlightGetRequests = new Map();

  const genKey = (url, config = {}) => {
    const params = config.params ? JSON.stringify(config.params) : '{}';
    return `${url}?${params}`;
  };

  instance.get = (url, config = {}) => {
    if (config?.disableDuplicate) {
      return originalGet(url, config);
    }

    const key = genKey(url, config);
    if (inFlightGetRequests.has(key)) {
      return inFlightGetRequests.get(key);
    }

    const reqPromise = originalGet(url, config).finally(() => {
      inFlightGetRequests.delete(key);
    });

    inFlightGetRequests.set(key, reqPromise);
    return reqPromise;
  };
}

patchAPIInstance(API);

export function updateAPI() {
  API = axios.create({
    baseURL: import.meta.env.VITE_REACT_APP_SERVER_URL
      ? import.meta.env.VITE_REACT_APP_SERVER_URL
      : '',
    headers: {
      'MedusaXD-API-User': getUserIdFromLocalStorage(),
      'Cache-Control': 'no-store',
    },
  });

  patchAPIInstance(API);
}

API.interceptors.response.use(
  (response) => response,
  (error) => {
    showError(error);
  },
);

// playground

// Build API request payload
export const buildApiPayload = (messages, systemPrompt, inputs, parameterEnabled) => {
  const processedMessages = messages
    .filter(isValidMessage)
    .map(formatMessageForAPI)
    .filter(Boolean);

  // If there's a system prompt, insert it at the beginning of messages
  if (systemPrompt && systemPrompt.trim()) {
    processedMessages.unshift({
      role: MESSAGE_ROLES.SYSTEM,
      content: systemPrompt.trim()
    });
  }

  const payload = {
    model: inputs.model,
    messages: processedMessages,
    group: inputs.group, 
    stream: inputs.stream,
  };

  // Add enabled parameters
  const parameterMappings = {
    temperature: 'temperature',
    top_p: 'top_p',
    max_tokens: 'max_tokens',
    frequency_penalty: 'frequency_penalty',
    presence_penalty: 'presence_penalty',
    seed: 'seed'
  };

  Object.entries(parameterMappings).forEach(([key, param]) => {
    if (parameterEnabled[key] && inputs[param] !== undefined && inputs[param] !== null) {
      payload[param] = inputs[param];
    }
  });

  return payload;
};

// Handle API error response
export const handleApiError = (error, response = null) => {
  const errorInfo = {
    error: error.message || 'Unknown error',
    timestamp: new Date().toISOString(),
    stack: error.stack
  };

  if (response) {
    errorInfo.status = response.status;
    errorInfo.statusText = response.statusText;
  }

  if (error.message.includes('HTTP error')) {
    errorInfo.details = 'Server returned error status code';
  } else if (error.message.includes('Failed to fetch')) {
    errorInfo.details = 'Network connection failed or server not responding';
  }

  return errorInfo;
};

// Handle model data
export const processModelsData = (data, currentModel) => {
  const modelOptions = data.map(model => ({
    label: model,
    value: model,
  }));

  const hasCurrentModel = modelOptions.some(option => option.value === currentModel);
  const selectedModel = hasCurrentModel && modelOptions.length > 0
    ? currentModel
    : modelOptions[0]?.value;

  return { modelOptions, selectedModel };
};

// Handle group data
export const processGroupsData = (data, userGroup) => {
  let groupOptions = Object.entries(data).map(([group, info]) => ({
    label: info.desc.length > 20 ? info.desc.substring(0, 20) + '...' : info.desc,
    value: group,
    ratio: info.ratio,
    fullLabel: info.desc,
  }));

  if (groupOptions.length === 0) {
    groupOptions = [{
      label: 'User Groups',
      value: '',
      ratio: 1,
    }];
  } else if (userGroup) {
    const userGroupIndex = groupOptions.findIndex(g => g.value === userGroup);
    if (userGroupIndex > -1) {
      const userGroupOption = groupOptions.splice(userGroupIndex, 1)[0];
      groupOptions.unshift(userGroupOption);
    }
  }

  return groupOptions;
};

// Originally utils.js from components

export async function getOAuthState() {
  let path = '/api/oauth/state';
  let affCode = localStorage.getItem('aff');
  if (affCode && affCode.length > 0) {
    path += `?aff=${affCode}`;
  }
  const res = await API.get(path);
  const { success, message, data } = res.data;
  if (success) {
    return data;
  } else {
    showError(message);
    return '';
  }
}

export async function onOIDCClicked(auth_url, client_id, openInNewTab = false) {
  const state = await getOAuthState();
  if (!state) return;
  const redirect_uri = `${window.location.origin}/oauth/oidc`;
  const response_type = 'code';
  const scope = 'openid profile email';
  const url = `${auth_url}?client_id=${client_id}&redirect_uri=${redirect_uri}&response_type=${response_type}&scope=${scope}&state=${state}`;
  if (openInNewTab) {
    window.open(url);
  } else {
    window.location.href = url;
  }
}

export async function onGitHubOAuthClicked(github_client_id) {
  const state = await getOAuthState();
  if (!state) return;
  window.open(
    `https://github.com/login/oauth/authorize?client_id=${github_client_id}&state=${state}&scope=user:email`,
  );
}

export async function onLinuxDOOAuthClicked(linuxdo_client_id) {
  const state = await getOAuthState();
  if (!state) return;
  window.open(
    `https://connect.linux.do/oauth2/authorize?response_type=code&client_id=${linuxdo_client_id}&state=${state}`,
  );
}

let channelModels = undefined;
export async function loadChannelModels() {
  const res = await API.get('/api/models');
  const { success, data } = res.data;
  if (!success) {
    return;
  }
  channelModels = data;
  localStorage.setItem('channel_models', JSON.stringify(data));
}

export function getChannelModels(type) {
  if (channelModels !== undefined && type in channelModels) {
    if (!channelModels[type]) {
      return [];
    }
    return channelModels[type];
  }
  let models = localStorage.getItem('channel_models');
  if (!models) {
    return [];
  }
  channelModels = JSON.parse(models);
  if (type in channelModels) {
    return channelModels[type];
  }
  return [];
}
