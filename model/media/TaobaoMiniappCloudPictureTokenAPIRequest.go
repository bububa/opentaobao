package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappcloudpicturetokenAPIRequest 云存储上传token获取 API请求
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
type TaobaominiappcloudpicturetokenAPIRequest struct {
	model.Params
	// 请求参数
	_generateTokenRequest *GenerateTokenRequest
}

// NewTaobaominiappcloudpicturetokenRequest 初始化TaobaominiappcloudpicturetokenAPIRequest对象
func NewTaobaominiappcloudpicturetokenRequest() *TaobaominiappcloudpicturetokenAPIRequest {
	return &TaobaominiappcloudpicturetokenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappcloudpicturetokenAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.picture.token"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappcloudpicturetokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappcloudpicturetokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGenerateTokenRequest is GenerateTokenRequest Setter
// 请求参数
func (r *TaobaominiappcloudpicturetokenAPIRequest) SetGenerateTokenRequest(_generateTokenRequest *GenerateTokenRequest) error {
	r._generateTokenRequest = _generateTokenRequest
	r.Set("generate_token_request", _generateTokenRequest)
	return nil
}

// GetGenerateTokenRequest GenerateTokenRequest Getter
func (r TaobaominiappcloudpicturetokenAPIRequest) GetGenerateTokenRequest() *GenerateTokenRequest {
	return r._generateTokenRequest
}
