package minimin

// HandleMinFile 配置文件转换到运行的结构体
func (d *DNA) HandleMinFile(file MinFile) (err error) {
	d.Env = file.Env
	for _, geneObj := range file.Genes {
		var newGeneObj GeneForWork
		newGeneObj.Status = Queue
		newGeneObj.Name = geneObj.Name
		newGeneObj.Ignore = geneObj.Ignore
		newGeneObj.When = geneObj.When
		newGeneObj.MPs = HandleMPs(geneObj.MPs)
		d.Genes = append(d.Genes, newGeneObj)
	}
	return err
}
