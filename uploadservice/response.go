package uploadservice

type CreateRequest struct {
	FileName string `json:"filename"` // 文件名称
	MD5      string `json:"md5"`      // 文件MD5值
	Size     uint64 `json:size`       // 文件大小，核验收到文件大小是否正常
	Slices   uint   `json:slices`     // 文件切片数
}

type CreateReponse struct {
	TaskId string `json:"taskid"` //为上传任务分配的唯一标识

}
