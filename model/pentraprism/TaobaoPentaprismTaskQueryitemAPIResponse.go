package pentraprism

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPentaprismTaskQueryitemAPIResponse
查询任务当前进度 API返回值
taobao.pentaprism.task.queryitem

外网用户查询五棱镜任务系统当前进度 */
type TaobaoPentaprismTaskQueryitemAPIResponse struct {
	model.CommonResponse
	TaobaoPentaprismTaskQueryitemAPIResponseModel
}

// TaobaoPentaprismTaskQueryitemAPIResponseModel is 查询任务当前进度 成功返回结果
type TaobaoPentaprismTaskQueryitemAPIResponseModel struct {
	XMLName xml.Name `xml:"pentaprism_task_queryitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// TOP接口标准出参
	Result *TaskResult `json:"result,omitempty" xml:"result,omitempty"`
}
