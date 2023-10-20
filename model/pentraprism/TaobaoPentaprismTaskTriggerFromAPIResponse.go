package pentraprism

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopentaprismtasktriggerfromAPIResponse 任务进度推进（根据fromtoken） API返回值
// taobao.pentaprism.task.trigger.from
//
// 外网用户推进单条五棱镜任务进度
type TaobaopentaprismtasktriggerfromAPIResponse struct {
	model.CommonResponse
	TaobaopentaprismtasktriggerfromAPIResponseModel
}

// TaobaopentaprismtasktriggerfromAPIResponseModel is 任务进度推进（根据fromtoken） 成功返回结果
type TaobaopentaprismtasktriggerfromAPIResponseModel struct {
	XMLName xml.Name `xml:"pentaprism_task_trigger_from_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// TOP接口标准出参
	Result *TaskResult `json:"result,omitempty" xml:"result,omitempty"`
}
