package plugins

import (
	"context"
	"fmt"
	"k8s.io/klog"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	frameworkruntime "k8s.io/kubernetes/pkg/scheduler/framework/runtime"

)

const Name = "sample-plugin"

var (
	_ framework.FilterPlugin    = &Sample{}
	_ framework.PreFilterPlugin = &Sample{}
	_ framework.PreBindPlugin = &Sample{}

	scheme = runtime.NewScheme()
)

type Args struct {
	FavoriteColor  string `json:"favorite_color,omitempty"`
	FavoriteNumber int    `json:"favorite_number,omitempty"`
	ThanksTo       string `json:"thanks_to,omitempty"`
}

type Sample struct {
	args   *Args
	handle framework.Handle
}

func New(rargs runtime.Object, h framework.Handle) (framework.Plugin, error)  {
	args := &Args{}
	if err := frameworkruntime.DecodeInto(rargs, args); err != nil {
		return nil, err
	}
	klog.V(3).Infof("get plugin config args: %+v", args)
	return &Sample{
		args: args,
		handle: h,
	}, nil
}

func (s *Sample) PreFilter(ctx context.Context, state *framework.CycleState, p *v1.Pod) *framework.Status {
	klog.V(3).Infof("prefilter pod: %v", p.Name)
	return framework.NewStatus(framework.Success, "")
}

func (s *Sample) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, nodeInfo.Node().GetName())
	return framework.NewStatus(framework.Success, "")
}

func (s *Sample) PreBind(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) *framework.Status {
	if nodeInfo, err := s.handle.SnapshotSharedLister().NodeInfos().Get(nodeName); err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("prebind get node info error: %+v", nodeName))
	}else {
		klog.V(3).Infof("prebind node info: %+v", nodeInfo.Node())
		return framework.NewStatus(framework.Success, "")
	}

}

func (s *Sample) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}

func (s *Sample) Name() string {
	return Name
}


