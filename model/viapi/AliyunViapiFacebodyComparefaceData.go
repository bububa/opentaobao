package viapi

// AliyunviapifacebodycomparefaceData 结构体
type AliyunviapifacebodycomparefaceData struct {
	// 误识率在10e-3,10e-4,10e-5时对应的置信度分类阈值
	Thresholds []int64 `json:"thresholds,omitempty" xml:"thresholds>int64,omitempty"`
	// 图片1最大人脸矩形框(left, top, width, height)，如图片中没有人脸，返回矩形框数值均为0
	RectAList []int64 `json:"rect_a_list,omitempty" xml:"rect_a_list>int64,omitempty"`
	// 图片2最大人脸矩形框(left, top, width, height)，如图片中没有人脸，返回矩形框数值均为0
	RectBList []int64 `json:"rect_b_list,omitempty" xml:"rect_b_list>int64,omitempty"`
	// 两张图片中最大人脸属于同一个人的置信度：0-100，如某张图片中没有人脸，返回置信度为0
	Confidence string `json:"confidence,omitempty" xml:"confidence,omitempty"`
}
