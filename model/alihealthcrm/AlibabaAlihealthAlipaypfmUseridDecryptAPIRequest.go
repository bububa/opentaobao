package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest 用户标识解密接口 API请求
// alibaba.alihealth.alipaypfm.userid.decrypt
//
// 用户唯一表示加密传输，调用方解密
type AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest struct {
	model.Params
	// 小程序appid
	_appId string
	// 加密后的userId
	_content string
}

// NewAlibabaAlihealthAlipaypfmUseridDecryptRequest 初始化AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest对象
func NewAlibabaAlihealthAlipaypfmUseridDecryptRequest() *AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest {
	return &AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.alipaypfm.userid.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAppId is AppId Setter
// 小程序appid
func (r *AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest) GetAppId() string {
	return r._appId
}

// SetContent is Content Setter
// 加密后的userId
func (r *AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest) GetContent() string {
	return r._content
}
