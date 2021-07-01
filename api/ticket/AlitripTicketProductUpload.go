package ticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ticket"
)

/* 
【门票API2.0】门票收费项目管理接口 
alitrip.ticket.product.upload

航旅度假新版门票商品（门票收费项目）管理接口：支持门票商品的发布、编辑。如果在ali_product_id下没有发布过门票商品，则系统将判断为新发布商品，否则是编辑已有商品。可以通过辅助查询接口判断是否已在某个ali_product_id下发布过门票商品。
对应新发布商品的情况，有些参数是必填的，请仔细查看各字段说明。对于编辑商品的情况，ali_product_id和out_product_id至少需要填一个，其他参数都是可选，编辑情况支持增量更新（某个参数不传则使用该商品上原有值）。
*/
func AlitripTicketProductUpload(clt *core.SDKClient, req *ticket.AlitripTicketProductUploadAPIRequest, session string) (*ticket.AlitripTicketProductUploadAPIResponse, error) {
    var resp ticket.AlitripTicketProductUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
