package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoGradesGetAPIRequest
分销商等级查询 API请求
taobao.fenxiao.grades.get

根据供应商ID，查询他的分销商等级信息 */
type TaobaoFenxiaoGradesGetAPIRequest struct {
	model.Params
}

// New
