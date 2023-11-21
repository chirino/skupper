package control_api

// Informer creates a *ApiListDevicesInOrganizationInformer which provides a simpler
// API to list devices but which is implemented with the Watch api.  The *ApiListDevicesInOrganizationInformer
// maintains a local device cache which gets updated with the Watch events.
func (r ApiListDevicesInVPCRequest) Informer() *Informer[ModelsDevice] {
	informer := NewInformer[ModelsDevice](&DeviceAdaptor{}, r.gtRevision, ApiWatchEventsRequest{
		ctx:        r.ctx,
		ApiService: r.ApiService.client.VPCApi,
		id:         r.id,
	})
	return informer
}

type DeviceAdaptor struct{}

func (d DeviceAdaptor) Revision(item ModelsDevice) int32 {
	return item.GetRevision()
}

func (d DeviceAdaptor) Key(item ModelsDevice) string {
	return item.GetId()
}

func (d DeviceAdaptor) Kind() string {
	return "device"
}

func (d DeviceAdaptor) Item(value map[string]interface{}) (ModelsDevice, error) {
	item := ModelsDevice{}
	err := JsonUnmarshal(value, &item)
	return item, err
}

var _ InformerAdaptor[ModelsDevice] = &DeviceAdaptor{}
