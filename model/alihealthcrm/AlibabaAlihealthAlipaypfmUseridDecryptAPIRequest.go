package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest
用户标识解密接口 API请求
alibaba.alihealth.alipaypfm.userid.decrypt

用户唯一表示加密传输，调用方解密 */
type AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest struct {
	model.Params
	// 小程序appid
	_appId string
	// 加密后的userId
	_content string
}

// New
