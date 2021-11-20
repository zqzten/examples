package v2

import (
	"fmt"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/zqzten/examples/conversion-webhook/api/v1"
)

func (src *NetworkAddress) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.NetworkAddress)
	dst.ObjectMeta = src.ObjectMeta

	parts := strings.Split(src.Spec.Address, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid address value: %s", src.Spec.Address)
	}
	switch src.Spec.HostType {
	case "ip":
		dst.Spec.IP = parts[0]
	case "domain":
		dst.Spec.Domain = parts[0]
	default:
		return fmt.Errorf("unrecognized host type: %s", src.Spec.HostType)
	}
	dst.Spec.Port = parts[1]

	dst.Spec.Note = src.Spec.Note

	return nil
}

func (dst *NetworkAddress) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.NetworkAddress)
	dst.ObjectMeta = src.ObjectMeta

	if len(src.Spec.IP) > 0 {
		dst.Spec.Address = fmt.Sprintf("%s:%s", src.Spec.IP, src.Spec.Port)
		dst.Spec.HostType = "ip"
	} else {
		dst.Spec.Address = fmt.Sprintf("%s:%s", src.Spec.Domain, src.Spec.Port)
		dst.Spec.HostType = "domain"
	}

	dst.Spec.Note = src.Spec.Note

	return nil
}
