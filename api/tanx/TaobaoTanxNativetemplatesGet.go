package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// Taobaotanxnativetemplatesget 批量获取本地模板信息
// taobao.tanx.nativetemplates.get
//
// 根据传入的本地模板ID批量返回本地模板
func Taobaotanxnativetemplatesget(clt *core.SDKClient, req *tanx.TaobaotanxnativetemplatesgetAPIRequest, session string) (*tanx.TaobaotanxnativetemplatesgetAPIResponse, error) {
	var resp tanx.TaobaotanxnativetemplatesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
