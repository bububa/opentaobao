package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Alibabaitempublishpropsget 商品级联属性信息获取
// alibaba.item.publish.props.get
//
// 新商品发布，商品级联属性信息获取
func Alibabaitempublishpropsget(clt *core.SDKClient, req *tbitem.AlibabaitempublishpropsgetAPIRequest, session string) (*tbitem.AlibabaitempublishpropsgetAPIResponse, error) {
	var resp tbitem.AlibabaitempublishpropsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
