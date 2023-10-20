package media

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudPictureTokenAPIRequest 云存储上传token获取 API请求
// taobao.miniapp.cloud.picture.token
//
// 获取图片上传token，作为http://upload.cloud.tmall.com/api/proxy/upload接口上传的凭证。
//
// 上传的html示例代码为：
// &amp;lt;form action=&amp;quot;http://upload.cloud.tmall.com/api/proxy/upload&amp;quot; method=&amp;quot;post&amp;quot; enctype=&amp;quot;multipart/form-data&amp;quot;&amp;gt;
//
//	上传token: &amp;lt;input type=&amp;quot;text&amp;quot; name=&amp;quot;Authorization&amp;quot; value=&amp;quot;&amp;quot; /&amp;gt;&amp;lt;br/&amp;gt;
//	上传文件:&amp;lt;input type=&amp;quot;file&amp;quot; name=&amp;quot;content&amp;quot; /&amp;gt;&amp;lt;br/&amp;gt;
//	&amp;lt;input type=&amp;quot;submit&amp;quot; value=&amp;quot;Submit &amp;quot;/&amp;gt;
//
// &amp;lt;/form&amp;gt;
type TaobaoMiniappCloudPictureTokenAPIRequest struct {
	model.Params
	// 请求参数
	_generateTokenRequest *GenerateTokenRequest
}

// NewTaobaoMiniappCloudPictureTokenRequest 初始化TaobaoMiniappCloudPictureTokenAPIRequest对象
func NewTaobaoMiniappCloudPictureTokenRequest() *TaobaoMiniappCloudPictureTokenAPIRequest {
	return &TaobaoMiniappCloudPictureTokenAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappCloudPictureTokenAPIRequest) Reset() {
	r._generateTokenRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudPictureTokenAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.picture.token"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudPictureTokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappCloudPictureTokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGenerateTokenRequest is GenerateTokenRequest Setter
// 请求参数
func (r *TaobaoMiniappCloudPictureTokenAPIRequest) SetGenerateTokenRequest(_generateTokenRequest *GenerateTokenRequest) error {
	r._generateTokenRequest = _generateTokenRequest
	r.Set("generate_token_request", _generateTokenRequest)
	return nil
}

// GetGenerateTokenRequest GenerateTokenRequest Getter
func (r TaobaoMiniappCloudPictureTokenAPIRequest) GetGenerateTokenRequest() *GenerateTokenRequest {
	return r._generateTokenRequest
}

var poolTaobaoMiniappCloudPictureTokenAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappCloudPictureTokenRequest()
	},
}

// GetTaobaoMiniappCloudPictureTokenRequest 从 sync.Pool 获取 TaobaoMiniappCloudPictureTokenAPIRequest
func GetTaobaoMiniappCloudPictureTokenAPIRequest() *TaobaoMiniappCloudPictureTokenAPIRequest {
	return poolTaobaoMiniappCloudPictureTokenAPIRequest.Get().(*TaobaoMiniappCloudPictureTokenAPIRequest)
}

// ReleaseTaobaoMiniappCloudPictureTokenAPIRequest 将 TaobaoMiniappCloudPictureTokenAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappCloudPictureTokenAPIRequest(v *TaobaoMiniappCloudPictureTokenAPIRequest) {
	v.Reset()
	poolTaobaoMiniappCloudPictureTokenAPIRequest.Put(v)
}
