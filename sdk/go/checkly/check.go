// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Check struct {
	pulumi.CustomResourceState

	Activated                 pulumi.BoolOutput                        `pulumi:"activated"`
	AlertChannelSubscriptions CheckAlertChannelSubscriptionArrayOutput `pulumi:"alertChannelSubscriptions"`
	AlertSettings             CheckAlertSettingsOutput                 `pulumi:"alertSettings"`
	DegradedResponseTime      pulumi.IntPtrOutput                      `pulumi:"degradedResponseTime"`
	DoubleCheck               pulumi.BoolPtrOutput                     `pulumi:"doubleCheck"`
	EnvironmentVariables      pulumi.MapOutput                         `pulumi:"environmentVariables"`
	Frequency                 pulumi.IntOutput                         `pulumi:"frequency"`
	FrequencyOffset           pulumi.IntPtrOutput                      `pulumi:"frequencyOffset"`
	GroupId                   pulumi.IntPtrOutput                      `pulumi:"groupId"`
	GroupOrder                pulumi.IntPtrOutput                      `pulumi:"groupOrder"`
	LocalSetupScript          pulumi.StringPtrOutput                   `pulumi:"localSetupScript"`
	LocalTeardownScript       pulumi.StringPtrOutput                   `pulumi:"localTeardownScript"`
	Locations                 pulumi.StringArrayOutput                 `pulumi:"locations"`
	MaxResponseTime           pulumi.IntPtrOutput                      `pulumi:"maxResponseTime"`
	Muted                     pulumi.BoolPtrOutput                     `pulumi:"muted"`
	Name                      pulumi.StringOutput                      `pulumi:"name"`
	Request                   CheckRequestPtrOutput                    `pulumi:"request"`
	RuntimeId                 pulumi.StringPtrOutput                   `pulumi:"runtimeId"`
	Script                    pulumi.StringPtrOutput                   `pulumi:"script"`
	SetupSnippetId            pulumi.IntPtrOutput                      `pulumi:"setupSnippetId"`
	ShouldFail                pulumi.BoolPtrOutput                     `pulumi:"shouldFail"`
	SslCheck                  pulumi.BoolPtrOutput                     `pulumi:"sslCheck"`
	Tags                      pulumi.StringArrayOutput                 `pulumi:"tags"`
	TeardownSnippetId         pulumi.IntPtrOutput                      `pulumi:"teardownSnippetId"`
	Type                      pulumi.StringOutput                      `pulumi:"type"`
	UseGlobalAlertSettings    pulumi.BoolPtrOutput                     `pulumi:"useGlobalAlertSettings"`
}

// NewCheck registers a new resource with the given unique name, arguments, and options.
func NewCheck(ctx *pulumi.Context,
	name string, args *CheckArgs, opts ...pulumi.ResourceOption) (*Check, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Activated == nil {
		return nil, errors.New("invalid value for required argument 'Activated'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Check
	err := ctx.RegisterResource("checkly:index/check:Check", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCheck gets an existing Check resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CheckState, opts ...pulumi.ResourceOption) (*Check, error) {
	var resource Check
	err := ctx.ReadResource("checkly:index/check:Check", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Check resources.
type checkState struct {
	Activated                 *bool                           `pulumi:"activated"`
	AlertChannelSubscriptions []CheckAlertChannelSubscription `pulumi:"alertChannelSubscriptions"`
	AlertSettings             *CheckAlertSettings             `pulumi:"alertSettings"`
	DegradedResponseTime      *int                            `pulumi:"degradedResponseTime"`
	DoubleCheck               *bool                           `pulumi:"doubleCheck"`
	EnvironmentVariables      map[string]interface{}          `pulumi:"environmentVariables"`
	Frequency                 *int                            `pulumi:"frequency"`
	FrequencyOffset           *int                            `pulumi:"frequencyOffset"`
	GroupId                   *int                            `pulumi:"groupId"`
	GroupOrder                *int                            `pulumi:"groupOrder"`
	LocalSetupScript          *string                         `pulumi:"localSetupScript"`
	LocalTeardownScript       *string                         `pulumi:"localTeardownScript"`
	Locations                 []string                        `pulumi:"locations"`
	MaxResponseTime           *int                            `pulumi:"maxResponseTime"`
	Muted                     *bool                           `pulumi:"muted"`
	Name                      *string                         `pulumi:"name"`
	Request                   *CheckRequest                   `pulumi:"request"`
	RuntimeId                 *string                         `pulumi:"runtimeId"`
	Script                    *string                         `pulumi:"script"`
	SetupSnippetId            *int                            `pulumi:"setupSnippetId"`
	ShouldFail                *bool                           `pulumi:"shouldFail"`
	SslCheck                  *bool                           `pulumi:"sslCheck"`
	Tags                      []string                        `pulumi:"tags"`
	TeardownSnippetId         *int                            `pulumi:"teardownSnippetId"`
	Type                      *string                         `pulumi:"type"`
	UseGlobalAlertSettings    *bool                           `pulumi:"useGlobalAlertSettings"`
}

type CheckState struct {
	Activated                 pulumi.BoolPtrInput
	AlertChannelSubscriptions CheckAlertChannelSubscriptionArrayInput
	AlertSettings             CheckAlertSettingsPtrInput
	DegradedResponseTime      pulumi.IntPtrInput
	DoubleCheck               pulumi.BoolPtrInput
	EnvironmentVariables      pulumi.MapInput
	Frequency                 pulumi.IntPtrInput
	FrequencyOffset           pulumi.IntPtrInput
	GroupId                   pulumi.IntPtrInput
	GroupOrder                pulumi.IntPtrInput
	LocalSetupScript          pulumi.StringPtrInput
	LocalTeardownScript       pulumi.StringPtrInput
	Locations                 pulumi.StringArrayInput
	MaxResponseTime           pulumi.IntPtrInput
	Muted                     pulumi.BoolPtrInput
	Name                      pulumi.StringPtrInput
	Request                   CheckRequestPtrInput
	RuntimeId                 pulumi.StringPtrInput
	Script                    pulumi.StringPtrInput
	SetupSnippetId            pulumi.IntPtrInput
	ShouldFail                pulumi.BoolPtrInput
	SslCheck                  pulumi.BoolPtrInput
	Tags                      pulumi.StringArrayInput
	TeardownSnippetId         pulumi.IntPtrInput
	Type                      pulumi.StringPtrInput
	UseGlobalAlertSettings    pulumi.BoolPtrInput
}

func (CheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*checkState)(nil)).Elem()
}

type checkArgs struct {
	Activated                 bool                            `pulumi:"activated"`
	AlertChannelSubscriptions []CheckAlertChannelSubscription `pulumi:"alertChannelSubscriptions"`
	AlertSettings             *CheckAlertSettings             `pulumi:"alertSettings"`
	DegradedResponseTime      *int                            `pulumi:"degradedResponseTime"`
	DoubleCheck               *bool                           `pulumi:"doubleCheck"`
	EnvironmentVariables      map[string]interface{}          `pulumi:"environmentVariables"`
	Frequency                 int                             `pulumi:"frequency"`
	FrequencyOffset           *int                            `pulumi:"frequencyOffset"`
	GroupId                   *int                            `pulumi:"groupId"`
	GroupOrder                *int                            `pulumi:"groupOrder"`
	LocalSetupScript          *string                         `pulumi:"localSetupScript"`
	LocalTeardownScript       *string                         `pulumi:"localTeardownScript"`
	Locations                 []string                        `pulumi:"locations"`
	MaxResponseTime           *int                            `pulumi:"maxResponseTime"`
	Muted                     *bool                           `pulumi:"muted"`
	Name                      *string                         `pulumi:"name"`
	Request                   *CheckRequest                   `pulumi:"request"`
	RuntimeId                 *string                         `pulumi:"runtimeId"`
	Script                    *string                         `pulumi:"script"`
	SetupSnippetId            *int                            `pulumi:"setupSnippetId"`
	ShouldFail                *bool                           `pulumi:"shouldFail"`
	SslCheck                  *bool                           `pulumi:"sslCheck"`
	Tags                      []string                        `pulumi:"tags"`
	TeardownSnippetId         *int                            `pulumi:"teardownSnippetId"`
	Type                      string                          `pulumi:"type"`
	UseGlobalAlertSettings    *bool                           `pulumi:"useGlobalAlertSettings"`
}

// The set of arguments for constructing a Check resource.
type CheckArgs struct {
	Activated                 pulumi.BoolInput
	AlertChannelSubscriptions CheckAlertChannelSubscriptionArrayInput
	AlertSettings             CheckAlertSettingsPtrInput
	DegradedResponseTime      pulumi.IntPtrInput
	DoubleCheck               pulumi.BoolPtrInput
	EnvironmentVariables      pulumi.MapInput
	Frequency                 pulumi.IntInput
	FrequencyOffset           pulumi.IntPtrInput
	GroupId                   pulumi.IntPtrInput
	GroupOrder                pulumi.IntPtrInput
	LocalSetupScript          pulumi.StringPtrInput
	LocalTeardownScript       pulumi.StringPtrInput
	Locations                 pulumi.StringArrayInput
	MaxResponseTime           pulumi.IntPtrInput
	Muted                     pulumi.BoolPtrInput
	Name                      pulumi.StringPtrInput
	Request                   CheckRequestPtrInput
	RuntimeId                 pulumi.StringPtrInput
	Script                    pulumi.StringPtrInput
	SetupSnippetId            pulumi.IntPtrInput
	ShouldFail                pulumi.BoolPtrInput
	SslCheck                  pulumi.BoolPtrInput
	Tags                      pulumi.StringArrayInput
	TeardownSnippetId         pulumi.IntPtrInput
	Type                      pulumi.StringInput
	UseGlobalAlertSettings    pulumi.BoolPtrInput
}

func (CheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*checkArgs)(nil)).Elem()
}

type CheckInput interface {
	pulumi.Input

	ToCheckOutput() CheckOutput
	ToCheckOutputWithContext(ctx context.Context) CheckOutput
}

func (*Check) ElementType() reflect.Type {
	return reflect.TypeOf((**Check)(nil)).Elem()
}

func (i *Check) ToCheckOutput() CheckOutput {
	return i.ToCheckOutputWithContext(context.Background())
}

func (i *Check) ToCheckOutputWithContext(ctx context.Context) CheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckOutput)
}

// CheckArrayInput is an input type that accepts CheckArray and CheckArrayOutput values.
// You can construct a concrete instance of `CheckArrayInput` via:
//
//          CheckArray{ CheckArgs{...} }
type CheckArrayInput interface {
	pulumi.Input

	ToCheckArrayOutput() CheckArrayOutput
	ToCheckArrayOutputWithContext(context.Context) CheckArrayOutput
}

type CheckArray []CheckInput

func (CheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Check)(nil)).Elem()
}

func (i CheckArray) ToCheckArrayOutput() CheckArrayOutput {
	return i.ToCheckArrayOutputWithContext(context.Background())
}

func (i CheckArray) ToCheckArrayOutputWithContext(ctx context.Context) CheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckArrayOutput)
}

// CheckMapInput is an input type that accepts CheckMap and CheckMapOutput values.
// You can construct a concrete instance of `CheckMapInput` via:
//
//          CheckMap{ "key": CheckArgs{...} }
type CheckMapInput interface {
	pulumi.Input

	ToCheckMapOutput() CheckMapOutput
	ToCheckMapOutputWithContext(context.Context) CheckMapOutput
}

type CheckMap map[string]CheckInput

func (CheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Check)(nil)).Elem()
}

func (i CheckMap) ToCheckMapOutput() CheckMapOutput {
	return i.ToCheckMapOutputWithContext(context.Background())
}

func (i CheckMap) ToCheckMapOutputWithContext(ctx context.Context) CheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckMapOutput)
}

type CheckOutput struct{ *pulumi.OutputState }

func (CheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Check)(nil)).Elem()
}

func (o CheckOutput) ToCheckOutput() CheckOutput {
	return o
}

func (o CheckOutput) ToCheckOutputWithContext(ctx context.Context) CheckOutput {
	return o
}

type CheckArrayOutput struct{ *pulumi.OutputState }

func (CheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Check)(nil)).Elem()
}

func (o CheckArrayOutput) ToCheckArrayOutput() CheckArrayOutput {
	return o
}

func (o CheckArrayOutput) ToCheckArrayOutputWithContext(ctx context.Context) CheckArrayOutput {
	return o
}

func (o CheckArrayOutput) Index(i pulumi.IntInput) CheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Check {
		return vs[0].([]*Check)[vs[1].(int)]
	}).(CheckOutput)
}

type CheckMapOutput struct{ *pulumi.OutputState }

func (CheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Check)(nil)).Elem()
}

func (o CheckMapOutput) ToCheckMapOutput() CheckMapOutput {
	return o
}

func (o CheckMapOutput) ToCheckMapOutputWithContext(ctx context.Context) CheckMapOutput {
	return o
}

func (o CheckMapOutput) MapIndex(k pulumi.StringInput) CheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Check {
		return vs[0].(map[string]*Check)[vs[1].(string)]
	}).(CheckOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CheckInput)(nil)).Elem(), &Check{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckArrayInput)(nil)).Elem(), CheckArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckMapInput)(nil)).Elem(), CheckMap{})
	pulumi.RegisterOutputType(CheckOutput{})
	pulumi.RegisterOutputType(CheckArrayOutput{})
	pulumi.RegisterOutputType(CheckMapOutput{})
}
