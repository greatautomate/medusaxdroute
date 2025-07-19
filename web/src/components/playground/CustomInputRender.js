import React from 'react';

const CustomInputRender = (props) => {
  const { detailProps } = props;
  const { clearContextNode, uploadNode, inputNode, sendNode, onClick } = detailProps;

  // Clear button
  const styledClearNode = clearContextNode
    ? React.cloneElement(clearContextNode, {
      className: `!rounded-full !bg-gray-100 hover:!bg-red-500 hover:!text-white flex-shrink-0 transition-all ${clearContextNode.props.className || ''}`,
      style: {
        ...clearContextNode.props.style,
        width: '32px',
        height: '32px',
        minWidth: '32px',
        padding: 0,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }
    })
    : null;

  // Send button
  const styledSendNode = React.cloneElement(sendNode, {
    className: `!rounded-full !bg-purple-500 hover:!bg-purple-600 flex-shrink-0 transition-all ${sendNode.props.className || ''}`,
    style: {
      ...sendNode.props.style,
      width: '32px',
      height: '32px',
      minWidth: '32px',
      padding: 0,
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    }
  });

  return (
    <div className="p-2 sm:p-4">
      <div
        className="flex items-center gap-2 sm:gap-3 p-2 bg-gray-50 rounded-xl sm:rounded-2xl shadow-sm hover:shadow-md transition-shadow"
        style={{ border: '1px solid var(--semi-color-border)' }}
        onClick={onClick}
      >
        {/* Clear conversation button - left */}
        {styledClearNode}
        <div className="flex-1">
          {inputNode}
        </div>
        {/* Send button - right */}
        {styledSendNode}
      </div>
    </div>
  );
};

export default CustomInputRender; 