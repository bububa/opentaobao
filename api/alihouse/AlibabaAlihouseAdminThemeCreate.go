package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseadminthemecreate 创建云主题
// alibaba.alihouse.admin.theme.create
//
// 创建云主题
func Alibabaalihouseadminthemecreate(clt *core.SDKClient, req *alihouse.AlibabaalihouseadminthemecreateAPIRequest, session string) (*alihouse.AlibabaalihouseadminthemecreateAPIResponse, error) {
	var resp alihouse.AlibabaalihouseadminthemecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
