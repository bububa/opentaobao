package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

/* TaobaoXhotelRateRelationshipwithrpGet
根据gid查询卖家下所有的rpId
taobao.xhotel.rate.relationshipwithrp.get

根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据 */
func TaobaoXhotelRateRelationshipwithrpGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateRelationshipwithrpGetAPIRequest, session string) (*xhotelitem.TaobaoXhotelRateRelationshipwithrpGetAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelRateRelationshipwithrpGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
