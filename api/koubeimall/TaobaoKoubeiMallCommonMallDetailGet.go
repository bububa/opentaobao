package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// Taobaokoubeimallcommonmalldetailget 查询商圈详细信息
// taobao.koubei.mall.common.mall.detail.get
//
// 查询口碑综合体-商圈详细信息，包含商圈基础信息、门店类目分类、商圈推荐商品等模块信息
func Taobaokoubeimallcommonmalldetailget(clt *core.SDKClient, req *koubeimall.TaobaokoubeimallcommonmalldetailgetAPIRequest, session string) (*koubeimall.TaobaokoubeimallcommonmalldetailgetAPIResponse, error) {
	var resp koubeimall.TaobaokoubeimallcommonmalldetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
