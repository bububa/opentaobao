package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthAlipaypfmUseridDecrypt 用户标识解密接口
// alibaba.alihealth.alipaypfm.userid.decrypt
//
// 用户唯一表示加密传输，调用方解密
func AlibabaAlihealthAlipaypfmUseridDecrypt(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthAlipaypfmUseridDecryptAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthAlipaypfmUseridDecryptAPIResponse, error) {
	var resp alihealthcrm.AlibabaAlihealthAlipaypfmUseridDecryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
