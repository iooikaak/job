package enum

type Gateway string

const (
	ServiceName Gateway = "job"
	JobYaml     Gateway = "config/job.yaml"
	JobJson     Gateway = "config/job.json"
)

func (g Gateway) String() string {
	return string(g)
}
