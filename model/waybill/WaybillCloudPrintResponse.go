package waybill

// WaybillCloudPrintResponse 
type WaybillCloudPrintResponse struct {

    // 请求id
    ObjectId   string `json:"object_id,omitempty"`

    // 面单号, 子母件模式下为子面单号
    WaybillCode   string `json:"waybill_code,omitempty"`

    // 云打印内容（encryptedData表示加密结果，data表示非加密结果）;模板内容,具体解释见<a href="http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.8cf9Nj&treeId=17&articleId=105085&docType=1#12">链接</a>
    PrintData   string `json:"print_data,omitempty"`

    // 子母件中的母单号，当为子母件模式时，需要此单号为实际挂载物流详情的单号，需要使用此单号进行发货，查询物流详情，非子母件，此字段为空
    ParentWaybillCode   string `json:"parent_waybill_code,omitempty"`

}
