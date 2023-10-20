package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtsmscodesend 喵零发送短信
// tmall.nrt.sms.code.send
//
// 喵零发送短信
func Tmallnrtsmscodesend(clt *core.SDKClient, req *nrt.TmallnrtsmscodesendAPIRequest, session string) (*nrt.TmallnrtsmscodesendAPIResponse, error) {
	var resp nrt.TmallnrtsmscodesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
