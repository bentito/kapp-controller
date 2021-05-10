// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	kappctrlv1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	packages "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apiserver/apis/packages"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AppTemplateSpec)(nil), (*packages.AppTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AppTemplateSpec_To_packages_AppTemplateSpec(a.(*AppTemplateSpec), b.(*packages.AppTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*packages.AppTemplateSpec)(nil), (*AppTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_packages_AppTemplateSpec_To_v1alpha1_AppTemplateSpec(a.(*packages.AppTemplateSpec), b.(*AppTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Maintainer)(nil), (*packages.Maintainer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Maintainer_To_packages_Maintainer(a.(*Maintainer), b.(*packages.Maintainer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*packages.Maintainer)(nil), (*Maintainer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_packages_Maintainer_To_v1alpha1_Maintainer(a.(*packages.Maintainer), b.(*Maintainer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Package)(nil), (*packages.Package)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Package_To_packages_Package(a.(*Package), b.(*packages.Package), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*packages.Package)(nil), (*Package)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_packages_Package_To_v1alpha1_Package(a.(*packages.Package), b.(*Package), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageList)(nil), (*packages.PackageList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageList_To_packages_PackageList(a.(*PackageList), b.(*packages.PackageList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*packages.PackageList)(nil), (*PackageList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_packages_PackageList_To_v1alpha1_PackageList(a.(*packages.PackageList), b.(*PackageList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageSpec)(nil), (*packages.PackageSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageSpec_To_packages_PackageSpec(a.(*PackageSpec), b.(*packages.PackageSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*packages.PackageSpec)(nil), (*PackageSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_packages_PackageSpec_To_v1alpha1_PackageSpec(a.(*packages.PackageSpec), b.(*PackageSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageVersion)(nil), (*packages.PackageVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageVersion_To_packages_PackageVersion(a.(*PackageVersion), b.(*packages.PackageVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*packages.PackageVersion)(nil), (*PackageVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_packages_PackageVersion_To_v1alpha1_PackageVersion(a.(*packages.PackageVersion), b.(*PackageVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageVersionList)(nil), (*packages.PackageVersionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageVersionList_To_packages_PackageVersionList(a.(*PackageVersionList), b.(*packages.PackageVersionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*packages.PackageVersionList)(nil), (*PackageVersionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_packages_PackageVersionList_To_v1alpha1_PackageVersionList(a.(*packages.PackageVersionList), b.(*PackageVersionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageVersionSpec)(nil), (*packages.PackageVersionSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageVersionSpec_To_packages_PackageVersionSpec(a.(*PackageVersionSpec), b.(*packages.PackageVersionSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*packages.PackageVersionSpec)(nil), (*PackageVersionSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_packages_PackageVersionSpec_To_v1alpha1_PackageVersionSpec(a.(*packages.PackageVersionSpec), b.(*PackageVersionSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ValuesSchema)(nil), (*packages.ValuesSchema)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ValuesSchema_To_packages_ValuesSchema(a.(*ValuesSchema), b.(*packages.ValuesSchema), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*packages.ValuesSchema)(nil), (*ValuesSchema)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_packages_ValuesSchema_To_v1alpha1_ValuesSchema(a.(*packages.ValuesSchema), b.(*ValuesSchema), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AppTemplateSpec_To_packages_AppTemplateSpec(in *AppTemplateSpec, out *packages.AppTemplateSpec, s conversion.Scope) error {
	out.Spec = (*kappctrlv1alpha1.AppSpec)(unsafe.Pointer(in.Spec))
	return nil
}

// Convert_v1alpha1_AppTemplateSpec_To_packages_AppTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha1_AppTemplateSpec_To_packages_AppTemplateSpec(in *AppTemplateSpec, out *packages.AppTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_AppTemplateSpec_To_packages_AppTemplateSpec(in, out, s)
}

func autoConvert_packages_AppTemplateSpec_To_v1alpha1_AppTemplateSpec(in *packages.AppTemplateSpec, out *AppTemplateSpec, s conversion.Scope) error {
	out.Spec = (*kappctrlv1alpha1.AppSpec)(unsafe.Pointer(in.Spec))
	return nil
}

// Convert_packages_AppTemplateSpec_To_v1alpha1_AppTemplateSpec is an autogenerated conversion function.
func Convert_packages_AppTemplateSpec_To_v1alpha1_AppTemplateSpec(in *packages.AppTemplateSpec, out *AppTemplateSpec, s conversion.Scope) error {
	return autoConvert_packages_AppTemplateSpec_To_v1alpha1_AppTemplateSpec(in, out, s)
}

func autoConvert_v1alpha1_Maintainer_To_packages_Maintainer(in *Maintainer, out *packages.Maintainer, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_Maintainer_To_packages_Maintainer is an autogenerated conversion function.
func Convert_v1alpha1_Maintainer_To_packages_Maintainer(in *Maintainer, out *packages.Maintainer, s conversion.Scope) error {
	return autoConvert_v1alpha1_Maintainer_To_packages_Maintainer(in, out, s)
}

func autoConvert_packages_Maintainer_To_v1alpha1_Maintainer(in *packages.Maintainer, out *Maintainer, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_packages_Maintainer_To_v1alpha1_Maintainer is an autogenerated conversion function.
func Convert_packages_Maintainer_To_v1alpha1_Maintainer(in *packages.Maintainer, out *Maintainer, s conversion.Scope) error {
	return autoConvert_packages_Maintainer_To_v1alpha1_Maintainer(in, out, s)
}

func autoConvert_v1alpha1_Package_To_packages_Package(in *Package, out *packages.Package, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PackageSpec_To_packages_PackageSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Package_To_packages_Package is an autogenerated conversion function.
func Convert_v1alpha1_Package_To_packages_Package(in *Package, out *packages.Package, s conversion.Scope) error {
	return autoConvert_v1alpha1_Package_To_packages_Package(in, out, s)
}

func autoConvert_packages_Package_To_v1alpha1_Package(in *packages.Package, out *Package, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_packages_PackageSpec_To_v1alpha1_PackageSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_packages_Package_To_v1alpha1_Package is an autogenerated conversion function.
func Convert_packages_Package_To_v1alpha1_Package(in *packages.Package, out *Package, s conversion.Scope) error {
	return autoConvert_packages_Package_To_v1alpha1_Package(in, out, s)
}

func autoConvert_v1alpha1_PackageList_To_packages_PackageList(in *PackageList, out *packages.PackageList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]packages.Package)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PackageList_To_packages_PackageList is an autogenerated conversion function.
func Convert_v1alpha1_PackageList_To_packages_PackageList(in *PackageList, out *packages.PackageList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageList_To_packages_PackageList(in, out, s)
}

func autoConvert_packages_PackageList_To_v1alpha1_PackageList(in *packages.PackageList, out *PackageList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Package)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_packages_PackageList_To_v1alpha1_PackageList is an autogenerated conversion function.
func Convert_packages_PackageList_To_v1alpha1_PackageList(in *packages.PackageList, out *PackageList, s conversion.Scope) error {
	return autoConvert_packages_PackageList_To_v1alpha1_PackageList(in, out, s)
}

func autoConvert_v1alpha1_PackageSpec_To_packages_PackageSpec(in *PackageSpec, out *packages.PackageSpec, s conversion.Scope) error {
	out.DisplayName = in.DisplayName
	out.LongDescription = in.LongDescription
	out.ShortDescription = in.ShortDescription
	out.IconSVGBase64 = in.IconSVGBase64
	out.ProviderName = in.ProviderName
	out.Maintainers = *(*[]packages.Maintainer)(unsafe.Pointer(&in.Maintainers))
	out.Categories = *(*[]string)(unsafe.Pointer(&in.Categories))
	out.SupportDescription = in.SupportDescription
	return nil
}

// Convert_v1alpha1_PackageSpec_To_packages_PackageSpec is an autogenerated conversion function.
func Convert_v1alpha1_PackageSpec_To_packages_PackageSpec(in *PackageSpec, out *packages.PackageSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageSpec_To_packages_PackageSpec(in, out, s)
}

func autoConvert_packages_PackageSpec_To_v1alpha1_PackageSpec(in *packages.PackageSpec, out *PackageSpec, s conversion.Scope) error {
	out.DisplayName = in.DisplayName
	out.LongDescription = in.LongDescription
	out.ShortDescription = in.ShortDescription
	out.IconSVGBase64 = in.IconSVGBase64
	out.ProviderName = in.ProviderName
	out.Maintainers = *(*[]Maintainer)(unsafe.Pointer(&in.Maintainers))
	out.Categories = *(*[]string)(unsafe.Pointer(&in.Categories))
	out.SupportDescription = in.SupportDescription
	return nil
}

// Convert_packages_PackageSpec_To_v1alpha1_PackageSpec is an autogenerated conversion function.
func Convert_packages_PackageSpec_To_v1alpha1_PackageSpec(in *packages.PackageSpec, out *PackageSpec, s conversion.Scope) error {
	return autoConvert_packages_PackageSpec_To_v1alpha1_PackageSpec(in, out, s)
}

func autoConvert_v1alpha1_PackageVersion_To_packages_PackageVersion(in *PackageVersion, out *packages.PackageVersion, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PackageVersionSpec_To_packages_PackageVersionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PackageVersion_To_packages_PackageVersion is an autogenerated conversion function.
func Convert_v1alpha1_PackageVersion_To_packages_PackageVersion(in *PackageVersion, out *packages.PackageVersion, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageVersion_To_packages_PackageVersion(in, out, s)
}

func autoConvert_packages_PackageVersion_To_v1alpha1_PackageVersion(in *packages.PackageVersion, out *PackageVersion, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_packages_PackageVersionSpec_To_v1alpha1_PackageVersionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_packages_PackageVersion_To_v1alpha1_PackageVersion is an autogenerated conversion function.
func Convert_packages_PackageVersion_To_v1alpha1_PackageVersion(in *packages.PackageVersion, out *PackageVersion, s conversion.Scope) error {
	return autoConvert_packages_PackageVersion_To_v1alpha1_PackageVersion(in, out, s)
}

func autoConvert_v1alpha1_PackageVersionList_To_packages_PackageVersionList(in *PackageVersionList, out *packages.PackageVersionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]packages.PackageVersion)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PackageVersionList_To_packages_PackageVersionList is an autogenerated conversion function.
func Convert_v1alpha1_PackageVersionList_To_packages_PackageVersionList(in *PackageVersionList, out *packages.PackageVersionList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageVersionList_To_packages_PackageVersionList(in, out, s)
}

func autoConvert_packages_PackageVersionList_To_v1alpha1_PackageVersionList(in *packages.PackageVersionList, out *PackageVersionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]PackageVersion)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_packages_PackageVersionList_To_v1alpha1_PackageVersionList is an autogenerated conversion function.
func Convert_packages_PackageVersionList_To_v1alpha1_PackageVersionList(in *packages.PackageVersionList, out *PackageVersionList, s conversion.Scope) error {
	return autoConvert_packages_PackageVersionList_To_v1alpha1_PackageVersionList(in, out, s)
}

func autoConvert_v1alpha1_PackageVersionSpec_To_packages_PackageVersionSpec(in *PackageVersionSpec, out *packages.PackageVersionSpec, s conversion.Scope) error {
	out.PackageName = in.PackageName
	out.Version = in.Version
	out.Licenses = *(*[]string)(unsafe.Pointer(&in.Licenses))
	out.ReleasedAt = in.ReleasedAt
	out.CapactiyRequirementsDescription = in.CapactiyRequirementsDescription
	out.ReleaseNotes = in.ReleaseNotes
	if err := Convert_v1alpha1_AppTemplateSpec_To_packages_AppTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ValuesSchema_To_packages_ValuesSchema(&in.ValuesSchema, &out.ValuesSchema, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PackageVersionSpec_To_packages_PackageVersionSpec is an autogenerated conversion function.
func Convert_v1alpha1_PackageVersionSpec_To_packages_PackageVersionSpec(in *PackageVersionSpec, out *packages.PackageVersionSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageVersionSpec_To_packages_PackageVersionSpec(in, out, s)
}

func autoConvert_packages_PackageVersionSpec_To_v1alpha1_PackageVersionSpec(in *packages.PackageVersionSpec, out *PackageVersionSpec, s conversion.Scope) error {
	out.PackageName = in.PackageName
	out.Version = in.Version
	out.Licenses = *(*[]string)(unsafe.Pointer(&in.Licenses))
	out.ReleasedAt = in.ReleasedAt
	out.CapactiyRequirementsDescription = in.CapactiyRequirementsDescription
	out.ReleaseNotes = in.ReleaseNotes
	if err := Convert_packages_AppTemplateSpec_To_v1alpha1_AppTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := Convert_packages_ValuesSchema_To_v1alpha1_ValuesSchema(&in.ValuesSchema, &out.ValuesSchema, s); err != nil {
		return err
	}
	return nil
}

// Convert_packages_PackageVersionSpec_To_v1alpha1_PackageVersionSpec is an autogenerated conversion function.
func Convert_packages_PackageVersionSpec_To_v1alpha1_PackageVersionSpec(in *packages.PackageVersionSpec, out *PackageVersionSpec, s conversion.Scope) error {
	return autoConvert_packages_PackageVersionSpec_To_v1alpha1_PackageVersionSpec(in, out, s)
}

func autoConvert_v1alpha1_ValuesSchema_To_packages_ValuesSchema(in *ValuesSchema, out *packages.ValuesSchema, s conversion.Scope) error {
	out.OpenAPIv3 = in.OpenAPIv3
	return nil
}

// Convert_v1alpha1_ValuesSchema_To_packages_ValuesSchema is an autogenerated conversion function.
func Convert_v1alpha1_ValuesSchema_To_packages_ValuesSchema(in *ValuesSchema, out *packages.ValuesSchema, s conversion.Scope) error {
	return autoConvert_v1alpha1_ValuesSchema_To_packages_ValuesSchema(in, out, s)
}

func autoConvert_packages_ValuesSchema_To_v1alpha1_ValuesSchema(in *packages.ValuesSchema, out *ValuesSchema, s conversion.Scope) error {
	out.OpenAPIv3 = in.OpenAPIv3
	return nil
}

// Convert_packages_ValuesSchema_To_v1alpha1_ValuesSchema is an autogenerated conversion function.
func Convert_packages_ValuesSchema_To_v1alpha1_ValuesSchema(in *packages.ValuesSchema, out *ValuesSchema, s conversion.Scope) error {
	return autoConvert_packages_ValuesSchema_To_v1alpha1_ValuesSchema(in, out, s)
}
