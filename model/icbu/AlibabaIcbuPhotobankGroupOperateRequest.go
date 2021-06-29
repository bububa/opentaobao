package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片银行分组操作接口 API请求
alibaba.icbu.photobank.group.operate

修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名
*/
type AlibabaIcbuPhotobankGroupOperateRequest struct {
    model.Params
    // 图片分组操作请求对象
    _photoGroupOperationRequest   *PhotoGroupOperationRequest
}

// 初始化AlibabaIcbuPhotobankGroupOperateRequest对象
func NewAlibabaIcbuPhotobankGroupOperateRequest() *AlibabaIcbuPhotobankGroupOperateRequest{
    return &AlibabaIcbuPhotobankGroupOperateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuPhotobankGroupOperateRequest) GetApiMethodName() string {
    return "alibaba.icbu.photobank.group.operate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuPhotobankGroupOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PhotoGroupOperationRequest Setter
// 图片分组操作请求对象
func (r *AlibabaIcbuPhotobankGroupOperateRequest) SetPhotoGroupOperationRequest(_photoGroupOperationRequest *PhotoGroupOperationRequest) error {
    r._photoGroupOperationRequest = _photoGroupOperationRequest
    r.Set("photo_group_operation_request", _photoGroupOperationRequest)
    return nil
}

// PhotoGroupOperationRequest Getter
func (r AlibabaIcbuPhotobankGroupOperateRequest) GetPhotoGroupOperationRequest() *PhotoGroupOperationRequest {
    return r._photoGroupOperationRequest
}
