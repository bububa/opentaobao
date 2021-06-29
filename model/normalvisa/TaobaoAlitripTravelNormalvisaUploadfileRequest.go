package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传电子签证 APIRequest
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

func NewTaobaoAlitripTravelNormalvisaUploadfileRequest() *TaobaoAlitripTravelNormalvisaUploadfileRequest{
    return &TaobaoAlitripTravelNormalvisaUploadfileRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.uploadfile"
}

func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelNormalvisaUploadfileRequest) SetFileBytes(fileBytes []*model.File) error {
    r.fileBytes = fileBytes
    r.Set("file_bytes", fileBytes)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetFileBytes() []*model.File {
    return r.fileBytes
}

func (r *TaobaoAlitripTravelNormalvisaUploadfileRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetFileName() string {
    return r.fileName
}

func (r *TaobaoAlitripTravelNormalvisaUploadfileRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaUploadfileRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

