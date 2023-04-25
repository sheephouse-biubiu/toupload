package uploadservice

type CreateRequest struct {
	FileName string `json:"filename"` // 文件名称
	Total    uint64 `json:total`      // 文件大小，核验收到文件大小是否正常
	Slices   uint   `json:slices`     // 文件切片数
}
