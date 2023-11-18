package control_api

// Informer creates a *ApiListSecurityGroupsInformer which provides a simpler
// API to list devices but which is implemented with the Watch api.  The *ApiListSecurityGroupsInformer
// maintains a local device cache which gets updated with the Watch events.
func (r ApiListSecurityGroupsInVPCRequest) Informer() *Informer[ModelsSecurityGroup] {
	informer := NewInformer[ModelsSecurityGroup](&SecurityGroupAdaptor{}, r.gtRevision, ApiWatchEventsRequest{
		ctx:        r.ctx,
		ApiService: r.ApiService.client.VPCApi,
		id:         r.id,
	})
	return informer
}

type SecurityGroupAdaptor struct{}

func (d SecurityGroupAdaptor) Revision(item ModelsSecurityGroup) int32 {
	return item.GetRevision()
}

func (d SecurityGroupAdaptor) Key(item ModelsSecurityGroup) string {
	return item.GetId()
}

func (d SecurityGroupAdaptor) Kind() string {
	return "security-group"
}

func (d SecurityGroupAdaptor) Item(value map[string]interface{}) (ModelsSecurityGroup, error) {
	item := ModelsSecurityGroup{}
	err := JsonUnmarshal(value, &item)
	return item, err
}

var _ InformerAdaptor[ModelsSecurityGroup] = &SecurityGroupAdaptor{}
