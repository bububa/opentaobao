package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappCloudPictureTokenAPIResponse
云存储上传token获取 API返回值
taobao.miniapp.cloud.picture.token

获取图片上传token，作为http://upload.cloud.tmall.com/api/proxy/upload接口上传的凭证。

上传的html示例代码为：
&lt;form action=&quot;http://upload.cloud.tmall.com/api/proxy/upload&quot; method=&quot;post&quot; enctype=&quot;multipart/form-data&quot;&gt;
	上传token: &lt;input type=&quot;text&quot; name=&quot;Authorization&quot; value=&quot;&quot; /&gt;&lt;br/&gt;
	上传文件:&lt;input type=&quot;file&quot; name=&quot;content&quot; /&gt;&lt;br/&gt;
	&lt;input type=&quot;submit&quot; value=&quot;Submit &quot;/&gt;
&lt;/form&gt; */
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
