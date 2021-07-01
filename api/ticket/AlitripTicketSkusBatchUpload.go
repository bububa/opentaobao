package ticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ticket"
)

/* AlitripTicketSkusBatchUpload
【门票API2.0】门票价格库存同步接口（多票种批量更新）
alitrip.ticket.skus.batch.upload

飞猪度假新版门票商品价格库存同步接口（多票种批量更新）。
注1、一个票种下可以挂多个规则（规则id必须不一样，每个规则实际对应了一个sku），同一个规则可以在不同票种下使用。
注2、日历库存和区间库存门票，统一使用DateInventory结构。对于日历库存门票请上传每一天的价格库存；对于区间库存门票，建议只上传开始和结束日期的价格库存，也支持上传每天的价格库存，系统会自动进行聚合（取第一天的价格为区间价格，累计所有天的库存为区间库存）。
注3、该接口同时支持 新增某个规则的价格库存 和 更新现有规则的价格库存。如果不清楚是否已在某个规则下上传过价格库存，请使用alitrip.ticket.product.query接口进行查询。如果该规则在该票种下已经存在，则该接口会判断为是价格库存更新操作。 */
func AlitripTicketSkusBatchUpload(clt *core.SDKClient, req *ticket.AlitripTicketSkusBatchUploadAPIRequest, session string) (*ticket.AlitripTicketSkusBatchUploadAPIResponse, error) {
	var resp ticket.AlitripTicketSkusBatchUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
