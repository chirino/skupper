package control_api

// Informer creates a *ApiListSitesInOrganizationInformer which provides a simpler
// API to list sites but which is implemented with the Watch api.  The *ApiListSitesInOrganizationInformer
// maintains a local site cache which gets updated with the Watch events.
func (r ApiListSitesInVPCRequest) Informer() *Informer[ModelsSite] {
	informer := NewInformer[ModelsSite](&SiteAdaptor{}, r.gtRevision, ApiWatchEventsRequest{
		ctx:        r.ctx,
		ApiService: r.ApiService.client.VPCApi,
		id:         r.id,
	})
	return informer
}

type SiteAdaptor struct{}

func (d SiteAdaptor) Revision(item ModelsSite) int32 {
	return item.GetRevision()
}

func (d SiteAdaptor) Key(item ModelsSite) string {
	return item.GetId()
}

func (d SiteAdaptor) Kind() string {
	return "site"
}

func (d SiteAdaptor) Item(value map[string]interface{}) (ModelsSite, error) {
	item := ModelsSite{}
	err := JsonUnmarshal(value, &item)
	return item, err
}

var _ InformerAdaptor[ModelsSite] = &SiteAdaptor{}
