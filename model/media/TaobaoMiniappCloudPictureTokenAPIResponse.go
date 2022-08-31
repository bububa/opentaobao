package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudPictureTokenAPIResponse 云存储上传token获取 API返回值
// taobao.miniapp.cloud.picture.token
//
// 获取图片上传token，作为http://upload.cloud.tmall.com/api/proxy/upload接口上传的凭证。
//
// 上传的html示例代码为：
// &amp;lt;form action=&amp;quot;http://upload.cloud.tmall.com/api/proxy/upload&amp;quot; method=&amp;quot;post&amp;quot; enctype=&amp;quot;multipart/form-data&amp;quot;&amp;gt;
// 	上传token: &amp;lt;input type=&amp;quot;text&amp;quot; name=&amp;quot;Authorization&amp;quot; value=&amp;quot;&amp;quot; /&amp;gt;&amp;lt;br/&amp;gt;
// 	上传文件:&amp;lt;input type=&amp;quot;file&amp;quot; name=&amp;quot;content&amp;quot; /&amp;gt;&amp;lt;br/&amp;gt;
// 	&amp;lt;input type=&amp;quot;submit&amp;quot; value=&amp;quot;Submit &amp;quot;/&amp;gt;
// &amp;lt;/form&amp;gt;
type TaobaoMiniappCloudPictureTokenAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappCloudPictureTokenAPIResponseModel
}

// TaobaoMiniappCloudPictureTokenAPIResponseModel is 云存储上传token获取 成功返回结果
type TaobaoMiniappCloudPictureTokenAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_cloud_picture_token_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
}
