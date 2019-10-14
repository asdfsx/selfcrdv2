package v1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	validationutils "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var selfcrdv2log = logf.Log.WithName("selfcrdv2-resource")

func (r *SelfCRDV2) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

var _ webhook.Defaulter = &SelfCRDV2{}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-core-clsutar-ai-v1-selfcrdv2,mutating=false,failurePolicy=fail,groups=core.clustar.ai,resources=selfcrdv2s,versions=v1,name=selfcrdv2.core.clustar.ai

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *SelfCRDV2) Default() {
	selfcrdv2log.Info("default", "name", r.Name)

	if r.Spec.CustomID == "" {
		r.Spec.CustomID = "CustomID"
	}
	if r.Spec.Username == "" {
		r.Spec.Username = "Username"
	}
}

var _ webhook.Validator = &SelfCRDV2{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SelfCRDV2) ValidateCreate() error {
	selfcrdv2log.Info("validate create", "name", r.Name)

	return r.validateSelfCRDV2()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SelfCRDV2) ValidateUpdate(old runtime.Object) error {
	selfcrdv2log.Info("validate update", "name", r.Name)

	return r.validateSelfCRDV2()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SelfCRDV2) ValidateDelete() error {
	selfcrdv2log.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *SelfCRDV2) validateSelfCRDV2() error {
	var allErrs field.ErrorList
	if err := r.validateSelfCRDV2Name(); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := r.validateSelfCRDV2Spec(); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "core.clustar.ai", Kind: "SelfCRDV2"},
		r.Name, allErrs)
}

func (r *SelfCRDV2) validateSelfCRDV2Spec() *field.Error {
	// The field helpers from the kubernetes API machinery help us return nicely
	// structured validation errors.
	return nil
}

func (r *SelfCRDV2) validateSelfCRDV2Name() *field.Error {
	if len(r.ObjectMeta.Name) > validationutils.DNS1035LabelMaxLength-11 {
		// The job name length is 63 character like all Kubernetes objects
		// (which must fit in a DNS subdomain). The cronjob controller appends
		// a 11-character suffix to the cronjob (`-$TIMESTAMP`) when creating
		// a job. The job name length limit is 63 characters. Therefore cronjob
		// names must have length <= 63-11=52. If we don't validate this here,
		// then job creation will fail later.
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be no more than 52 characters")
	}
	return nil
}

// +kubebuilder:docs-gen:collapse=Validate object name
