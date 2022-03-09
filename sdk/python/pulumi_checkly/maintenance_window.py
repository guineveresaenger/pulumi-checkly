# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MaintenanceWindowArgs', 'MaintenanceWindow']

@pulumi.input_type
class MaintenanceWindowArgs:
    def __init__(__self__, *,
                 ends_at: pulumi.Input[str],
                 starts_at: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 repeat_ends_at: Optional[pulumi.Input[str]] = None,
                 repeat_interval: Optional[pulumi.Input[int]] = None,
                 repeat_unit: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a MaintenanceWindow resource.
        """
        pulumi.set(__self__, "ends_at", ends_at)
        pulumi.set(__self__, "starts_at", starts_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if repeat_ends_at is not None:
            pulumi.set(__self__, "repeat_ends_at", repeat_ends_at)
        if repeat_interval is not None:
            pulumi.set(__self__, "repeat_interval", repeat_interval)
        if repeat_unit is not None:
            pulumi.set(__self__, "repeat_unit", repeat_unit)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="endsAt")
    def ends_at(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ends_at")

    @ends_at.setter
    def ends_at(self, value: pulumi.Input[str]):
        pulumi.set(self, "ends_at", value)

    @property
    @pulumi.getter(name="startsAt")
    def starts_at(self) -> pulumi.Input[str]:
        return pulumi.get(self, "starts_at")

    @starts_at.setter
    def starts_at(self, value: pulumi.Input[str]):
        pulumi.set(self, "starts_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="repeatEndsAt")
    def repeat_ends_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repeat_ends_at")

    @repeat_ends_at.setter
    def repeat_ends_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repeat_ends_at", value)

    @property
    @pulumi.getter(name="repeatInterval")
    def repeat_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "repeat_interval")

    @repeat_interval.setter
    def repeat_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repeat_interval", value)

    @property
    @pulumi.getter(name="repeatUnit")
    def repeat_unit(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repeat_unit")

    @repeat_unit.setter
    def repeat_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repeat_unit", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _MaintenanceWindowState:
    def __init__(__self__, *,
                 ends_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repeat_ends_at: Optional[pulumi.Input[str]] = None,
                 repeat_interval: Optional[pulumi.Input[int]] = None,
                 repeat_unit: Optional[pulumi.Input[str]] = None,
                 starts_at: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering MaintenanceWindow resources.
        """
        if ends_at is not None:
            pulumi.set(__self__, "ends_at", ends_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if repeat_ends_at is not None:
            pulumi.set(__self__, "repeat_ends_at", repeat_ends_at)
        if repeat_interval is not None:
            pulumi.set(__self__, "repeat_interval", repeat_interval)
        if repeat_unit is not None:
            pulumi.set(__self__, "repeat_unit", repeat_unit)
        if starts_at is not None:
            pulumi.set(__self__, "starts_at", starts_at)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="endsAt")
    def ends_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ends_at")

    @ends_at.setter
    def ends_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ends_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="repeatEndsAt")
    def repeat_ends_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repeat_ends_at")

    @repeat_ends_at.setter
    def repeat_ends_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repeat_ends_at", value)

    @property
    @pulumi.getter(name="repeatInterval")
    def repeat_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "repeat_interval")

    @repeat_interval.setter
    def repeat_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repeat_interval", value)

    @property
    @pulumi.getter(name="repeatUnit")
    def repeat_unit(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repeat_unit")

    @repeat_unit.setter
    def repeat_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repeat_unit", value)

    @property
    @pulumi.getter(name="startsAt")
    def starts_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "starts_at")

    @starts_at.setter
    def starts_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "starts_at", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class MaintenanceWindow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ends_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repeat_ends_at: Optional[pulumi.Input[str]] = None,
                 repeat_interval: Optional[pulumi.Input[int]] = None,
                 repeat_unit: Optional[pulumi.Input[str]] = None,
                 starts_at: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a MaintenanceWindow resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MaintenanceWindowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a MaintenanceWindow resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param MaintenanceWindowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MaintenanceWindowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ends_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repeat_ends_at: Optional[pulumi.Input[str]] = None,
                 repeat_interval: Optional[pulumi.Input[int]] = None,
                 repeat_unit: Optional[pulumi.Input[str]] = None,
                 starts_at: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MaintenanceWindowArgs.__new__(MaintenanceWindowArgs)

            if ends_at is None and not opts.urn:
                raise TypeError("Missing required property 'ends_at'")
            __props__.__dict__["ends_at"] = ends_at
            __props__.__dict__["name"] = name
            __props__.__dict__["repeat_ends_at"] = repeat_ends_at
            __props__.__dict__["repeat_interval"] = repeat_interval
            __props__.__dict__["repeat_unit"] = repeat_unit
            if starts_at is None and not opts.urn:
                raise TypeError("Missing required property 'starts_at'")
            __props__.__dict__["starts_at"] = starts_at
            __props__.__dict__["tags"] = tags
        super(MaintenanceWindow, __self__).__init__(
            'checkly:index/maintenanceWindow:MaintenanceWindow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ends_at: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            repeat_ends_at: Optional[pulumi.Input[str]] = None,
            repeat_interval: Optional[pulumi.Input[int]] = None,
            repeat_unit: Optional[pulumi.Input[str]] = None,
            starts_at: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'MaintenanceWindow':
        """
        Get an existing MaintenanceWindow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MaintenanceWindowState.__new__(_MaintenanceWindowState)

        __props__.__dict__["ends_at"] = ends_at
        __props__.__dict__["name"] = name
        __props__.__dict__["repeat_ends_at"] = repeat_ends_at
        __props__.__dict__["repeat_interval"] = repeat_interval
        __props__.__dict__["repeat_unit"] = repeat_unit
        __props__.__dict__["starts_at"] = starts_at
        __props__.__dict__["tags"] = tags
        return MaintenanceWindow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endsAt")
    def ends_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ends_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="repeatEndsAt")
    def repeat_ends_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "repeat_ends_at")

    @property
    @pulumi.getter(name="repeatInterval")
    def repeat_interval(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "repeat_interval")

    @property
    @pulumi.getter(name="repeatUnit")
    def repeat_unit(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "repeat_unit")

    @property
    @pulumi.getter(name="startsAt")
    def starts_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "starts_at")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

