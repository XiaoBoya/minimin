package minimin

// MP 最小的单位，CICD最小的执行单位
type MP struct {
	Name        string            `yaml:"name" json:"name"`
	Description string            `yaml:"description,omitempty" json:"description,omitempty"`
	AminoAcid   string            `yaml:"amino_acid" json:"amino_acid"`
	Chain       string            `yaml:"chain,omitempty" json:"chain,omitempty"`
	Params      map[string]string `yaml:"params,omitempty" json:"params,omitempty"`
	Ignore      bool              `yaml:"ignore,omitempty" json:"ignore,omitempty"`
	When        string            `yaml:"when,omitempty" json:"when,omitempty"`
	Status      string            `yaml:"-" json:"status,omitempty"`
}

// MPForWork 支线
type MPForWork struct {
	Single *MP
	Multi  map[string][]MP
}

// Gene （基因）解析文件产生的对应顺序
type Gene struct {
	Name   string `yaml:"name" json:"name"`
	MPs    []MP   `yaml:"MPs" json:"MPs"`
	Ignore bool   `yaml:"ignore,omitempty" json:"ignore,omitempty"`
	When   string `yaml:"when,omitempty" json:"when,omitempty"`
}

// GeneForWork （基因）实际运行使用的阶段级别的数据结构
type GeneForWork struct {
	Name   string      `yaml:"name" json:"name"`
	MPs    []MPForWork `yaml:"MPs" json:"MPs"`
	Ignore bool        `yaml:"ignore,omitempty" json:"ignore,omitempty"`
	When   string      `yaml:"when,omitempty" json:"when,omitempty"`
	Status string      `yaml:"status" json:"status"`
}

// MinFile 配置文件结构
type MinFile struct {
	Name  string            `yaml:"name" json:"name"`
	Env   map[string]string `yaml:"env" json:"env"`
	Genes []Gene            `yaml:"genes" json:"genes"`
}

// DNA 运行使用的大数据结构
type DNA struct {
	Name  string            `yaml:"name" json:"name"`
	Env   map[string]string `yaml:"env" json:"env"`
	Genes []GeneForWork     `yaml:"genes" json:"genes"`
}
