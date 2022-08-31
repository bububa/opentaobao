package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeAdviserMessageNotice 催促小B发送短信
// alibaba.alihouse.newhome.adviser.message.notice
//
// 催促小B发送短信
func AlibabaAlihouseNewhomeAdviserMessageNotice(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
