// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows you to define alerting channels for the checks and groups in your account
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/checkly/pulumi-checkly/sdk/go/checkly"
// 	"github.com/pulumi/pulumi-checkly/sdk/go/checkly"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		emailAc, err := checkly.NewAlertChannel(ctx, "emailAc", &checkly.AlertChannelArgs{
// 			Email: &AlertChannelEmailArgs{
// 				Address: pulumi.String("john@example.com"),
// 			},
// 			SendRecovery:       pulumi.Bool(true),
// 			SendFailure:        pulumi.Bool(false),
// 			SendDegraded:       pulumi.Bool(true),
// 			SslExpiry:          pulumi.Bool(true),
// 			SslExpiryThreshold: pulumi.Int(22),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		smsAc, err := checkly.NewAlertChannel(ctx, "smsAc", &checkly.AlertChannelArgs{
// 			Sms: &AlertChannelSmsArgs{
// 				Name:   pulumi.String("john"),
// 				Number: pulumi.String("+5491100001111"),
// 			},
// 			SendRecovery: pulumi.Bool(true),
// 			SendFailure:  pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = checkly.NewAlertChannel(ctx, "slackAc", &checkly.AlertChannelArgs{
// 			Slack: &AlertChannelSlackArgs{
// 				Channel: pulumi.String("#checkly-notifications"),
// 				Url:     pulumi.String("https://slack.com/webhook"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = checkly.NewAlertChannel(ctx, "opsgenieAc", &checkly.AlertChannelArgs{
// 			Opsgenie: &AlertChannelOpsgenieArgs{
// 				Name:     pulumi.String("opsalerts"),
// 				ApiKey:   pulumi.String("fookey"),
// 				Region:   pulumi.String("fooregion"),
// 				Priority: pulumi.String("foopriority"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = checkly.NewAlertChannel(ctx, "pagerdutyAc", &checkly.AlertChannelArgs{
// 			Pagerduty: &AlertChannelPagerdutyArgs{
// 				Account:     pulumi.String("checkly"),
// 				ServiceKey:  pulumi.String("key1"),
// 				ServiceName: pulumi.String("pdalert"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = checkly.NewAlertChannel(ctx, "webhookAc", &checkly.AlertChannelArgs{
// 			Webhook: &AlertChannelWebhookArgs{
// 				Name:          pulumi.String("foo"),
// 				Method:        pulumi.String("get"),
// 				Template:      pulumi.String("footemplate"),
// 				Url:           pulumi.String("https://example.com/foo"),
// 				WebhookSecret: pulumi.String("foosecret"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = checkly.NewCheck(ctx, "example-check", &checkly.CheckArgs{
// 			AlertChannelSubscriptions: CheckAlertChannelSubscriptionArray{
// 				&CheckAlertChannelSubscriptionArgs{
// 					ChannelId: emailAc.ID(),
// 					Activated: pulumi.Bool(true),
// 				},
// 				&CheckAlertChannelSubscriptionArgs{
// 					ChannelId: smsAc.ID(),
// 					Activated: pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AlertChannel struct {
	pulumi.CustomResourceState

	Email     AlertChannelEmailPtrOutput     `pulumi:"email"`
	Opsgenie  AlertChannelOpsgeniePtrOutput  `pulumi:"opsgenie"`
	Pagerduty AlertChannelPagerdutyPtrOutput `pulumi:"pagerduty"`
	// (Default `false`)
	SendDegraded pulumi.BoolPtrOutput `pulumi:"sendDegraded"`
	// (Default `true`)
	SendFailure pulumi.BoolPtrOutput `pulumi:"sendFailure"`
	// (Default `true`)
	SendRecovery pulumi.BoolPtrOutput       `pulumi:"sendRecovery"`
	Slack        AlertChannelSlackPtrOutput `pulumi:"slack"`
	Sms          AlertChannelSmsPtrOutput   `pulumi:"sms"`
	// (Default `false`)
	SslExpiry pulumi.BoolPtrOutput `pulumi:"sslExpiry"`
	// Value must be between 1 and 30 (Default `30`)
	SslExpiryThreshold pulumi.IntPtrOutput          `pulumi:"sslExpiryThreshold"`
	Webhook            AlertChannelWebhookPtrOutput `pulumi:"webhook"`
}

// NewAlertChannel registers a new resource with the given unique name, arguments, and options.
func NewAlertChannel(ctx *pulumi.Context,
	name string, args *AlertChannelArgs, opts ...pulumi.ResourceOption) (*AlertChannel, error) {
	if args == nil {
		args = &AlertChannelArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AlertChannel
	err := ctx.RegisterResource("checkly:index/alertChannel:AlertChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertChannel gets an existing AlertChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertChannelState, opts ...pulumi.ResourceOption) (*AlertChannel, error) {
	var resource AlertChannel
	err := ctx.ReadResource("checkly:index/alertChannel:AlertChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertChannel resources.
type alertChannelState struct {
	Email     *AlertChannelEmail     `pulumi:"email"`
	Opsgenie  *AlertChannelOpsgenie  `pulumi:"opsgenie"`
	Pagerduty *AlertChannelPagerduty `pulumi:"pagerduty"`
	// (Default `false`)
	SendDegraded *bool `pulumi:"sendDegraded"`
	// (Default `true`)
	SendFailure *bool `pulumi:"sendFailure"`
	// (Default `true`)
	SendRecovery *bool              `pulumi:"sendRecovery"`
	Slack        *AlertChannelSlack `pulumi:"slack"`
	Sms          *AlertChannelSms   `pulumi:"sms"`
	// (Default `false`)
	SslExpiry *bool `pulumi:"sslExpiry"`
	// Value must be between 1 and 30 (Default `30`)
	SslExpiryThreshold *int                 `pulumi:"sslExpiryThreshold"`
	Webhook            *AlertChannelWebhook `pulumi:"webhook"`
}

type AlertChannelState struct {
	Email     AlertChannelEmailPtrInput
	Opsgenie  AlertChannelOpsgeniePtrInput
	Pagerduty AlertChannelPagerdutyPtrInput
	// (Default `false`)
	SendDegraded pulumi.BoolPtrInput
	// (Default `true`)
	SendFailure pulumi.BoolPtrInput
	// (Default `true`)
	SendRecovery pulumi.BoolPtrInput
	Slack        AlertChannelSlackPtrInput
	Sms          AlertChannelSmsPtrInput
	// (Default `false`)
	SslExpiry pulumi.BoolPtrInput
	// Value must be between 1 and 30 (Default `30`)
	SslExpiryThreshold pulumi.IntPtrInput
	Webhook            AlertChannelWebhookPtrInput
}

func (AlertChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertChannelState)(nil)).Elem()
}

type alertChannelArgs struct {
	Email     *AlertChannelEmail     `pulumi:"email"`
	Opsgenie  *AlertChannelOpsgenie  `pulumi:"opsgenie"`
	Pagerduty *AlertChannelPagerduty `pulumi:"pagerduty"`
	// (Default `false`)
	SendDegraded *bool `pulumi:"sendDegraded"`
	// (Default `true`)
	SendFailure *bool `pulumi:"sendFailure"`
	// (Default `true`)
	SendRecovery *bool              `pulumi:"sendRecovery"`
	Slack        *AlertChannelSlack `pulumi:"slack"`
	Sms          *AlertChannelSms   `pulumi:"sms"`
	// (Default `false`)
	SslExpiry *bool `pulumi:"sslExpiry"`
	// Value must be between 1 and 30 (Default `30`)
	SslExpiryThreshold *int                 `pulumi:"sslExpiryThreshold"`
	Webhook            *AlertChannelWebhook `pulumi:"webhook"`
}

// The set of arguments for constructing a AlertChannel resource.
type AlertChannelArgs struct {
	Email     AlertChannelEmailPtrInput
	Opsgenie  AlertChannelOpsgeniePtrInput
	Pagerduty AlertChannelPagerdutyPtrInput
	// (Default `false`)
	SendDegraded pulumi.BoolPtrInput
	// (Default `true`)
	SendFailure pulumi.BoolPtrInput
	// (Default `true`)
	SendRecovery pulumi.BoolPtrInput
	Slack        AlertChannelSlackPtrInput
	Sms          AlertChannelSmsPtrInput
	// (Default `false`)
	SslExpiry pulumi.BoolPtrInput
	// Value must be between 1 and 30 (Default `30`)
	SslExpiryThreshold pulumi.IntPtrInput
	Webhook            AlertChannelWebhookPtrInput
}

func (AlertChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertChannelArgs)(nil)).Elem()
}

type AlertChannelInput interface {
	pulumi.Input

	ToAlertChannelOutput() AlertChannelOutput
	ToAlertChannelOutputWithContext(ctx context.Context) AlertChannelOutput
}

func (*AlertChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertChannel)(nil)).Elem()
}

func (i *AlertChannel) ToAlertChannelOutput() AlertChannelOutput {
	return i.ToAlertChannelOutputWithContext(context.Background())
}

func (i *AlertChannel) ToAlertChannelOutputWithContext(ctx context.Context) AlertChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertChannelOutput)
}

// AlertChannelArrayInput is an input type that accepts AlertChannelArray and AlertChannelArrayOutput values.
// You can construct a concrete instance of `AlertChannelArrayInput` via:
//
//          AlertChannelArray{ AlertChannelArgs{...} }
type AlertChannelArrayInput interface {
	pulumi.Input

	ToAlertChannelArrayOutput() AlertChannelArrayOutput
	ToAlertChannelArrayOutputWithContext(context.Context) AlertChannelArrayOutput
}

type AlertChannelArray []AlertChannelInput

func (AlertChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertChannel)(nil)).Elem()
}

func (i AlertChannelArray) ToAlertChannelArrayOutput() AlertChannelArrayOutput {
	return i.ToAlertChannelArrayOutputWithContext(context.Background())
}

func (i AlertChannelArray) ToAlertChannelArrayOutputWithContext(ctx context.Context) AlertChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertChannelArrayOutput)
}

// AlertChannelMapInput is an input type that accepts AlertChannelMap and AlertChannelMapOutput values.
// You can construct a concrete instance of `AlertChannelMapInput` via:
//
//          AlertChannelMap{ "key": AlertChannelArgs{...} }
type AlertChannelMapInput interface {
	pulumi.Input

	ToAlertChannelMapOutput() AlertChannelMapOutput
	ToAlertChannelMapOutputWithContext(context.Context) AlertChannelMapOutput
}

type AlertChannelMap map[string]AlertChannelInput

func (AlertChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertChannel)(nil)).Elem()
}

func (i AlertChannelMap) ToAlertChannelMapOutput() AlertChannelMapOutput {
	return i.ToAlertChannelMapOutputWithContext(context.Background())
}

func (i AlertChannelMap) ToAlertChannelMapOutputWithContext(ctx context.Context) AlertChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertChannelMapOutput)
}

type AlertChannelOutput struct{ *pulumi.OutputState }

func (AlertChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertChannel)(nil)).Elem()
}

func (o AlertChannelOutput) ToAlertChannelOutput() AlertChannelOutput {
	return o
}

func (o AlertChannelOutput) ToAlertChannelOutputWithContext(ctx context.Context) AlertChannelOutput {
	return o
}

type AlertChannelArrayOutput struct{ *pulumi.OutputState }

func (AlertChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertChannel)(nil)).Elem()
}

func (o AlertChannelArrayOutput) ToAlertChannelArrayOutput() AlertChannelArrayOutput {
	return o
}

func (o AlertChannelArrayOutput) ToAlertChannelArrayOutputWithContext(ctx context.Context) AlertChannelArrayOutput {
	return o
}

func (o AlertChannelArrayOutput) Index(i pulumi.IntInput) AlertChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlertChannel {
		return vs[0].([]*AlertChannel)[vs[1].(int)]
	}).(AlertChannelOutput)
}

type AlertChannelMapOutput struct{ *pulumi.OutputState }

func (AlertChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertChannel)(nil)).Elem()
}

func (o AlertChannelMapOutput) ToAlertChannelMapOutput() AlertChannelMapOutput {
	return o
}

func (o AlertChannelMapOutput) ToAlertChannelMapOutputWithContext(ctx context.Context) AlertChannelMapOutput {
	return o
}

func (o AlertChannelMapOutput) MapIndex(k pulumi.StringInput) AlertChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlertChannel {
		return vs[0].(map[string]*AlertChannel)[vs[1].(string)]
	}).(AlertChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertChannelInput)(nil)).Elem(), &AlertChannel{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertChannelArrayInput)(nil)).Elem(), AlertChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertChannelMapInput)(nil)).Elem(), AlertChannelMap{})
	pulumi.RegisterOutputType(AlertChannelOutput{})
	pulumi.RegisterOutputType(AlertChannelArrayOutput{})
	pulumi.RegisterOutputType(AlertChannelMapOutput{})
}
