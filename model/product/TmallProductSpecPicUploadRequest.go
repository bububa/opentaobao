package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传产品规格认证图片 API请求
tmall.product.spec.pic.upload

上传指定类型的产品规格认证文件，并返回存有上传成功图片url的产品规格对象
*/
type TmallProductSpecPicUploadRequest struct {
    model.Params
    // 上传的认证图片的认证类型<br>1：代表产品包装正面图<br>2：代表完整产品资质<br>3：代表产品包装反面图<br>4：代表产品包装侧面图<br>5：代表产品包装条形码特写<br>6：代表特殊用途化妆品批准文号<br>7：代表3C认证图标<br>
    _certifyType   int64
    // 上传的认证图片文件
    _certifyPic   []*model.File
}

// 初始化TmallProductSpecPicUploadRequest对象
func NewTmallProductSpecPicUploadRequest() *TmallProductSpecPicUploadRequest{
    return &TmallProductSpecPicUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSpecPicUploadRequest) GetApiMethodName() string {
    return "tmall.product.spec.pic.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSpecPicUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CertifyType Setter
// 上传的认证图片的认证类型<br>1：代表产品包装正面图<br>2：代表完整产品资质<br>3：代表产品包装反面图<br>4：代表产品包装侧面图<br>5：代表产品包装条形码特写<br>6：代表特殊用途化妆品批准文号<br>7：代表3C认证图标<br>
func (r *TmallProductSpecPicUploadRequest) SetCertifyType(_certifyType int64) error {
    r._certifyType = _certifyType
    r.Set("certify_type", _certifyType)
    return nil
}

// CertifyType Getter
func (r TmallProductSpecPicUploadRequest) GetCertifyType() int64 {
    return r._certifyType
}
// CertifyPic Setter
// 上传的认证图片文件
func (r *TmallProductSpecPicUploadRequest) SetCertifyPic(_certifyPic []*model.File) error {
    r._certifyPic = _certifyPic
    r.Set("certify_pic", _certifyPic)
    return nil
}

// CertifyPic Getter
func (r TmallProductSpecPicUploadRequest) GetCertifyPic() []*model.File {
    return r._certifyPic
}
