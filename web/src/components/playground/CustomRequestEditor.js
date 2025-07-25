import React, { useState, useEffect } from 'react';
import {
  TextArea,
  Typography,
  Button,
  Switch,
  Banner,
} from '@douyinfe/semi-ui';
import {
  Code,
  Edit,
  Check,
  X,
  AlertTriangle,
} from 'lucide-react';
import { useTranslation } from 'react-i18next';

const CustomRequestEditor = ({
  customRequestMode,
  customRequestBody,
  onCustomRequestModeChange,
  onCustomRequestBodyChange,
  defaultPayload,
}) => {
  const { t } = useTranslation();
  const [isValid, setIsValid] = useState(true);
  const [errorMessage, setErrorMessage] = useState('');
  const [localValue, setLocalValue] = useState(customRequestBody || '');

  // Initialize with default payload when switching to custom mode
  useEffect(() => {
    if (customRequestMode && (!customRequestBody || customRequestBody.trim() === '')) {
      const defaultJson = defaultPayload ? JSON.stringify(defaultPayload, null, 2) : '';
      setLocalValue(defaultJson);
      onCustomRequestBodyChange(defaultJson);
    }
  }, [customRequestMode, defaultPayload, customRequestBody, onCustomRequestBodyChange]);

  // Sync external customRequestBody to local state
  useEffect(() => {
    if (customRequestBody !== localValue) {
      setLocalValue(customRequestBody || '');
      validateJson(customRequestBody || '');
    }
  }, [customRequestBody]);

  // Validate JSON format
  const validateJson = (value) => {
    if (!value.trim()) {
      setIsValid(true);
      setErrorMessage('');
      return true;
    }

    try {
      JSON.parse(value);
      setIsValid(true);
      setErrorMessage('');
      return true;
    } catch (error) {
      setIsValid(false);
      setErrorMessage(`JSON format error: ${error.message}`);
      return false;
    }
  };

  const handleValueChange = (value) => {
    setLocalValue(value);
    validateJson(value);
    // Always save user input, let preview logic handle JSON parsing errors
    onCustomRequestBodyChange(value);
  };

  const handleModeToggle = (enabled) => {
    onCustomRequestModeChange(enabled);
    if (enabled && defaultPayload) {
      const defaultJson = JSON.stringify(defaultPayload, null, 2);
      setLocalValue(defaultJson);
      onCustomRequestBodyChange(defaultJson);
    }
  };

  const formatJson = () => {
    try {
      const parsed = JSON.parse(localValue);
      const formatted = JSON.stringify(parsed, null, 2);
      setLocalValue(formatted);
      onCustomRequestBodyChange(formatted);
      setIsValid(true);
      setErrorMessage('');
    } catch (error) {
      // If formatting fails, keep as is
    }
  };

  return (
    <div className="space-y-4">
      {/* Custom mode toggle */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-2">
          <Code size={16} className="text-gray-500" />
          <Typography.Text strong className="text-sm">
            {t('自定义请求体模式')}
          </Typography.Text>
        </div>
        <Switch
          checked={customRequestMode}
          onChange={handleModeToggle}
          checkedText={t('开')}
          uncheckedText={t('关')}
          size="small"
        />
      </div>

      {customRequestMode && (
        <>
          {/* Warning message */}
          <Banner
            type="warning"
            description={t('启用此模式后，将使用您自定义的请求体发送API请求，模型配置面板的参数设置将被忽略。')}
            icon={<AlertTriangle size={16} />}
            className="!rounded-lg"
            closable={false}
          />

          {/* JSON editor */}
          <div>
            <div className="flex items-center justify-between mb-2">
              <Typography.Text strong className="text-sm">
                {t('请求体 JSON')}
              </Typography.Text>
              <div className="flex items-center gap-2">
                {isValid ? (
                  <div className="flex items-center gap-1 text-green-600">
                    <Check size={14} />
                    <Typography.Text className="text-xs">
                      {t('格式正确')}
                    </Typography.Text>
                  </div>
                ) : (
                  <div className="flex items-center gap-1 text-red-600">
                    <X size={14} />
                    <Typography.Text className="text-xs">
                      {t('格式错误')}
                    </Typography.Text>
                  </div>
                )}
                <Button
                  theme="borderless"
                  type="tertiary"
                  size="small"
                  icon={<Edit size={14} />}
                  onClick={formatJson}
                  disabled={!isValid}
                  className="!rounded-lg"
                >
                  {t('格式化')}
                </Button>
              </div>
            </div>

            <TextArea
              value={localValue}
              onChange={handleValueChange}
              placeholder='{"model": "gpt-4o", "messages": [...], ...}'
              autosize={{ minRows: 8, maxRows: 20 }}
              className={`custom-request-textarea !rounded-lg font-mono text-sm ${!isValid ? '!border-red-500' : ''}`}
              style={{
                fontFamily: 'Consolas, Monaco, "Courier New", monospace',
                lineHeight: '1.5',
              }}
            />

            {!isValid && errorMessage && (
              <Typography.Text type="danger" className="text-xs mt-1 block">
                {errorMessage}
              </Typography.Text>
            )}

            <Typography.Text className="text-xs text-gray-500 mt-2 block">
              {t('请输入有效的JSON格式的请求体。您可以参考预览面板中的默认请求体格式。')}
            </Typography.Text>
          </div>
        </>
      )}
    </div>
  );
};

export default CustomRequestEditor; 