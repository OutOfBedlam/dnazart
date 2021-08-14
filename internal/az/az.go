package az

type BuildList struct {
	Count int          `json:"count"`
	Value []BuildValue `json:"value"`
}

type BuildValue struct {
	Id                  int        `json:"id"`
	BuildNumber         string     `json:"buildNumber"`
	BuildNumberRevision int        `json:"buildNumberRevision"`
	Status              string     `json:"status"`
	Result              string     `json:"result"`
	QueueTime           string     `json:"queueTime"`
	StartTime           string     `json:"startTime"`
	FinishTime          string     `json:"finishTime"`
	Url                 string     `json:"url"`
	Uri                 string     `json:"uri"`
	SourceBranch        string     `json:"sourceBranch"`
	SourceVersion       string     `json:"sourceVersion"`
	Definition          Definition `json:"definition"`
	// Queue struct `json:"queue"`

	Project           Project            `json:"project"`
	Links             Links              `json:"_links"`
	Properties        map[string]string  `json:"properties"`
	Tags              []string           `json:"tags"`
	ValidationResults []ValidationResult `json:"validationResults"`
}

type Definition struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Url         string  `json:"url"`
	Uri         string  `json:"uri"`
	Path        string  `json:"path"`
	Type        string  `json:"type"`
	QueueStatus string  `json:"queueStatus"`
	Revision    int     `json:"revision"`
	Project     Project `json:"project"`
}
type Project struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Url            string `json:"url"`
	State          string `json:"state"`
	Revision       int    `json:"revision"`
	Visibility     string `json:"visibility"`
	LastUpdateTime string `json:"lastUpdateTime"`
}

type Links struct {
	Self                    Link `json:"self"`
	Web                     Link `json:"web"`
	SourceVersionDisplayUri Link `json:"sourceVersionDisplayUri"`
	Timeline                Link `json:"timeline"`
	Badge                   Link `json:"badge"`
}

type Link struct {
	Href string `json:"href"`
}

type ValidationResult struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

type ArtifactList struct {
	Count int        `json:"count"`
	Value []Artifact `json:"value"`
}

type Artifact struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Source   string   `json:"source"`
	Resource Resource `json:"resource"`
}

type Resource struct {
	Type        string            `json:"type"`
	Data        string            `json:"data"`
	Properties  map[string]string `json:"properties"`
	Url         string            `json:"url"`
	DownloadUrl string            `json:"downloadUrl"`
}
