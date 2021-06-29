package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传电子签证 API请求
taobao.alitrip.travel.normalvisa.uploadfile

上传电子签证
*/
type TaobaoAlitripTravelNormalvisaUploadfileRequest struct {
    model.Params
    // 文件
    fileBytes   []*model.File
    // 文件名：注意文件名请保证和上传的文件一直
    fileName   string
    // 订单id
    bizOrderId   int64
}

// 初始化TaobaoAlitripTravelNormalvisaUploadfileRequest对象
func NewTaobaoAlitripTravelNormalvisaUploadfileRequest() *TaobaoAlitripTravelNormalvisaUploadfileRequest{
    return &TaobaoAlitripTravelNormalvisaUploadfileRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.uploadfile"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FileBytes Setter
// 文件
func (r *TaobaoAlitripTravelNormalvisaUploadfileRequest) SetFileBytes(fileBytes []*model.File) error {
    r.fileBytes = fileBytes
    r.Set("file_bytes", fileBytes)
    return nil
}

// FileBytes Getter
func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetFileBytes() []*model.File {
    return r.fileBytes
}
// FileName Setter
// 文件名：注意文件名请保证和上传的文件一直
func (r *TaobaoAlitripTravelNormalvisaUploadfileRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

// FileName Getter
func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetFileName() string {
    return r.fileName
}
// BizOrderId Setter
// 订单id
func (r *TaobaoAlitripTravelNormalvisaUploadfileRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
