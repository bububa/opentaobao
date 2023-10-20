package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemhscodedetailget 通过hscode获取计量单位
// tmall.item.hscode.detail.get
//
// 通过hscode获取计量单位和销售单位
func Tmallitemhscodedetailget(clt *core.SDKClient, req *tbitem.TmallitemhscodedetailgetAPIRequest, session string) (*tbitem.TmallitemhscodedetailgetAPIResponse, error) {
	var resp tbitem.TmallitemhscodedetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
