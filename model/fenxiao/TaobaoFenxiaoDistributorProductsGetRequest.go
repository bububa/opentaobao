package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商查询产品信息 APIRequest
taobao.fenxiao.distributor.products.get

分销商查询供应商产品信息
*/
type TaobaoFenxiaoDistributorProductsGetRequest struct {
    model.Params

    // order_by
    orderBy   string 

    // time_type
    timeType   string 

    // 下载状态，默认是未下载；可选值：UNDOWNLOAD：未下载 ；DOWNLOADED：已下载；下载：指将供应商授权的产品发布为店铺新宝贝的过程，下载成功后，分销商需要将新生成的宝贝重新编辑并上架后售卖。
    downloadStatus   string 

    // 分销方式；可选择：AGENT ： 代销；DEALER：经销；DIRECT：直营
    tradeType   string 

    // 结束时间
    endTime   string 

    // 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
    fields   []String 

    // 根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
    itemIds   []Number 

    // 产品线ID
    productcatId   int64 

    // 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
    pids   []Number 

    // 开始修改时间
    startTime   string 

    // 页码（大于0的整数，默认1）
    pageNo   int64 

    // 每页记录数（默认20，最大50）
    pageSize   int64 

    // 供应商nick，分页查询时，必传
    supplierNick   string 

}

func NewTaobaoFenxiaoDistributorProductsGetRequest() *TaobaoFenxiaoDistributorProductsGetRequest{
    return &TaobaoFenxiaoDistributorProductsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.distributor.products.get"
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetOrderBy(orderBy string) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetOrderBy() string {
    return r.orderBy
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetTimeType(timeType string) error {
    r.timeType = timeType
    r.Set("time_type", timeType)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetTimeType() string {
    return r.timeType
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetDownloadStatus(downloadStatus string) error {
    r.downloadStatus = downloadStatus
    r.Set("download_status", downloadStatus)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetDownloadStatus() string {
    return r.downloadStatus
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetTradeType(tradeType string) error {
    r.tradeType = tradeType
    r.Set("trade_type", tradeType)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetTradeType() string {
    return r.tradeType
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetFields(fields []String) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetFields() []String {
    return r.fields
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetItemIds(itemIds []Number) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetItemIds() []Number {
    return r.itemIds
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetProductcatId(productcatId int64) error {
    r.productcatId = productcatId
    r.Set("productcat_id", productcatId)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetProductcatId() int64 {
    return r.productcatId
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetPids(pids []Number) error {
    r.pids = pids
    r.Set("pids", pids)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetPids() []Number {
    return r.pids
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetSupplierNick(supplierNick string) error {
    r.supplierNick = supplierNick
    r.Set("supplier_nick", supplierNick)
    return nil
}

func (r TaobaoFenxiaoDistributorProductsGetRequest) GetSupplierNick() string {
    return r.supplierNick
}

