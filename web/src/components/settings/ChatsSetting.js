import React, { useEffect, useState } from 'react';
import { Card, Spin } from '@douyinfe/semi-ui';
import SettingsChats from '../../pages/Setting/Chat/SettingsChats.js';
import { API, showError } from '../../helpers';
import { useTranslation } from 'react-i18next';

const ChatsSetting = () => {
  const { t } = useTranslation();
  let [inputs, setInputs] = useState({
    /* Chat settings */
    Chats: '[]',
  });

  let [loading, setLoading] = useState(false);

  const getOptions = async () => {
    const res = await API.get('/api/option/');
    const { success, message, data } = res.data;
    if (success) {
      let newInputs = {};
      data.forEach((item) => {
        if (
          item.key.endsWith('Enabled') ||
          ['DefaultCollapseSidebar'].includes(item.key)
        ) {
          newInputs[item.key] = item.value === 'true' ? true : false;
        } else {
          newInputs[item.key] = item.value;
        }
      });

      setInputs(newInputs);
    } else {
      showError(message);
    }
  };

  async function onRefresh() {
    try {
      setLoading(true);
      await getOptions();
    } catch (error) {
      showError(t('刷新失败'));
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    onRefresh();
  }, []);

  return (
    <>
      <Spin spinning={loading} size='large'>
        {/* Chat settings */}
        <Card style={{ marginTop: '10px' }}>
          <SettingsChats options={inputs} refresh={onRefresh} />
        </Card>
      </Spin>
    </>
  );
};

export default ChatsSetting; 