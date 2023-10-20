package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSpecPicUploadAPIRequest 上传产品规格认证图片 API请求
// tmall.product.spec.pic.upload
//
// 上传指定类型的产品规格认证文件，并返回存有上传成功图片url的产品规格对象
type TmallProductSpecPicUploadAPIRequest struct {
	model.Params
	// 上传的认证图片的认证类型<br>1：代表产品包装正面图<br>2：代表完整产品资质<br>3：代表产品包装反面图<br>4：代表产品包装侧面图<br>5：代表产品包装条形码特写<br>6：代表特殊用途化妆品批准文号<br>7：代表3C认证图标<br>
	_certifyType int64
	// 上传的认证图片文件
	_certifyPic *model.File
}

// NewTmallProductSpecPicUploadRequest 初始化TmallProductSpecPicUploadAPIRequest对象
func NewTmallProductSpecPicUploadRequest() *TmallProductSpecPicUploadAPIRequest {
	return &TmallProductSpecPicUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductSpecPicUploadAPIRequest) GetApiMethodName() string {
	return "tmall.product.spec.pic.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductSpecPicUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallProductSpecPicUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCertifyType is CertifyType Setter
// 上传的认证图片的认证类型&lt;br&gt;1：代表产品包装正面图&lt;br&gt;2：代表完整产品资质&lt;br&gt;3：代表产品包装反面图&lt;br&gt;4：代表产品包装侧面图&lt;br&gt;5：代表产品包装条形码特写&lt;br&gt;6：代表特殊用途化妆品批准文号&lt;br&gt;7：代表3C认证图标&lt;br&gt;
func (r *TmallProductSpecPicUploadAPIRequest) SetCertifyType(_certifyType int64) error {
	r._certifyType = _certifyType
	r.Set("certify_type", _certifyType)
	return nil
}

// GetCertifyType CertifyType Getter
func (r TmallProductSpecPicUploadAPIRequest) GetCertifyType() int64 {
	return r._certifyType
}

// SetCertifyPic is CertifyPic Setter
// 上传的认证图片文件
func (r *TmallProductSpecPicUploadAPIRequest) SetCertifyPic(_certifyPic *model.File) error {
	r._certifyPic = _certifyPic
	r.Set("certify_pic", _certifyPic)
	return nil
}

// GetCertifyPic CertifyPic Getter
func (r TmallProductSpecPicUploadAPIRequest) GetCertifyPic() *model.File {
	return r._certifyPic
}
