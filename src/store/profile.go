package store
//性能数据，主要是对于加载mmap的文件判断


type profile struct {
	HotDatas []string //热数据对应的文件名集合
	ColdDatas []string // 冷数据对应的文件名集合
}

func (p *profile) getFirstHot()string{

	return "1"
}
