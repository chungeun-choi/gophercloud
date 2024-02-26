package testing

import (
	"context"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/extensions/volumehost"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/extensions/volumetenants"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/v2/pagination"
	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
)

func TestListWithExtensions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	count := 0

	volumes.List(client.ServiceClient(), &volumes.ListOpts{}).EachPage(context.TODO(), func(_ context.Context, page pagination.Page) (bool, error) {
		count++
		actual, err := volumes.ExtractVolumes(page)
		if err != nil {
			t.Errorf("Failed to extract volumes: %v", err)
			return false, err
		}

		expected := []volumes.Volume{
			{
				ID:   "289da7f8-6440-407c-9fb4-7db01ec49164",
				Name: "vol-001",
				Attachments: []volumes.Attachment{{
					ServerID:     "83ec2e3b-4321-422b-8706-a84185f52a0a",
					AttachmentID: "05551600-a936-4d4a-ba42-79a037c1-c91a",
					AttachedAt:   time.Date(2016, 8, 6, 14, 48, 20, 0, time.UTC),
					HostName:     "foobar",
					VolumeID:     "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
					Device:       "/dev/vdc",
					ID:           "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
				}},
				AvailabilityZone:   "nova",
				Bootable:           "false",
				ConsistencyGroupID: "",
				CreatedAt:          time.Date(2015, 9, 17, 3, 35, 3, 0, time.UTC),
				Description:        "",
				Encrypted:          false,
				Metadata:           map[string]string{"foo": "bar"},
				Multiattach:        false,
				//TenantID:                  "304dc00909ac4d0da6c62d816bcb3459",
				//ReplicationDriverData:     "",
				//ReplicationExtendedStatus: "",
				ReplicationStatus: "disabled",
				Size:              75,
				SnapshotID:        "",
				SourceVolID:       "",
				Status:            "available",
				UserID:            "ff1ce52c03ab433aaba9108c2e3ef541",
				VolumeType:        "lvmdriver-1",
			},
			{
				ID:                 "96c3bda7-c82a-4f50-be73-ca7621794835",
				Name:               "vol-002",
				Attachments:        []volumes.Attachment{},
				AvailabilityZone:   "nova",
				Bootable:           "false",
				ConsistencyGroupID: "",
				CreatedAt:          time.Date(2015, 9, 17, 3, 32, 29, 0, time.UTC),
				Description:        "",
				Encrypted:          false,
				Metadata:           map[string]string{},
				Multiattach:        false,
				//TenantID:                  "304dc00909ac4d0da6c62d816bcb3459",
				//ReplicationDriverData:     "",
				//ReplicationExtendedStatus: "",
				ReplicationStatus: "disabled",
				Size:              75,
				SnapshotID:        "",
				SourceVolID:       "",
				Status:            "available",
				UserID:            "ff1ce52c03ab433aaba9108c2e3ef541",
				VolumeType:        "lvmdriver-1",
			},
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListAllWithExtensions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	type VolumeWithExt struct {
		volumes.Volume
		volumetenants.VolumeTenantExt
		volumehost.VolumeHostExt
	}

	allPages, err := volumes.List(client.ServiceClient(), &volumes.ListOpts{}).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	var actual []VolumeWithExt
	err = volumes.ExtractVolumesInto(allPages, &actual)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(actual))
	th.AssertEquals(t, "host-001", actual[0].Host)
	th.AssertEquals(t, "", actual[1].Host)
	th.AssertEquals(t, "304dc00909ac4d0da6c62d816bcb3459", actual[0].TenantID)
}

func TestListAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	allPages, err := volumes.List(client.ServiceClient(), &volumes.ListOpts{}).AllPages(context.TODO())
	th.AssertNoErr(t, err)
	actual, err := volumes.ExtractVolumes(allPages)
	th.AssertNoErr(t, err)

	expected := []volumes.Volume{
		{
			ID:   "289da7f8-6440-407c-9fb4-7db01ec49164",
			Name: "vol-001",
			Attachments: []volumes.Attachment{{
				ServerID:     "83ec2e3b-4321-422b-8706-a84185f52a0a",
				AttachmentID: "05551600-a936-4d4a-ba42-79a037c1-c91a",
				AttachedAt:   time.Date(2016, 8, 6, 14, 48, 20, 0, time.UTC),
				HostName:     "foobar",
				VolumeID:     "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
				Device:       "/dev/vdc",
				ID:           "d6cacb1a-8b59-4c88-ad90-d70ebb82bb75",
			}},
			AvailabilityZone:   "nova",
			Bootable:           "false",
			ConsistencyGroupID: "",
			CreatedAt:          time.Date(2015, 9, 17, 3, 35, 3, 0, time.UTC),
			Description:        "",
			Encrypted:          false,
			Metadata:           map[string]string{"foo": "bar"},
			Multiattach:        false,
			//TenantID:                  "304dc00909ac4d0da6c62d816bcb3459",
			//ReplicationDriverData:     "",
			//ReplicationExtendedStatus: "",
			ReplicationStatus: "disabled",
			Size:              75,
			SnapshotID:        "",
			SourceVolID:       "",
			Status:            "available",
			UserID:            "ff1ce52c03ab433aaba9108c2e3ef541",
			VolumeType:        "lvmdriver-1",
		},
		{
			ID:                 "96c3bda7-c82a-4f50-be73-ca7621794835",
			Name:               "vol-002",
			Attachments:        []volumes.Attachment{},
			AvailabilityZone:   "nova",
			Bootable:           "false",
			ConsistencyGroupID: "",
			CreatedAt:          time.Date(2015, 9, 17, 3, 32, 29, 0, time.UTC),
			Description:        "",
			Encrypted:          false,
			Metadata:           map[string]string{},
			Multiattach:        false,
			//TenantID:                  "304dc00909ac4d0da6c62d816bcb3459",
			//ReplicationDriverData:     "",
			//ReplicationExtendedStatus: "",
			ReplicationStatus: "disabled",
			Size:              75,
			SnapshotID:        "",
			SourceVolID:       "",
			Status:            "available",
			UserID:            "ff1ce52c03ab433aaba9108c2e3ef541",
			VolumeType:        "lvmdriver-1",
		},
	}

	th.CheckDeepEquals(t, expected, actual)

}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	v, err := volumes.Get(context.TODO(), client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v.Name, "vol-001")
	th.AssertEquals(t, v.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	options := &volumes.CreateOpts{Size: 75, Name: "vol-001"}
	n, err := volumes.Create(context.TODO(), client.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.Size, 75)
	th.AssertEquals(t, n.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteResponse(t)

	res := volumes.Delete(context.TODO(), client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22", volumes.DeleteOpts{})
	th.AssertNoErr(t, res.Err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateResponse(t)

	var name = "vol-002"
	options := volumes.UpdateOpts{Name: &name}
	v, err := volumes.Update(context.TODO(), client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22", options).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "vol-002", v.Name)
}

func TestGetWithExtensions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	var s struct {
		volumes.Volume
		volumetenants.VolumeTenantExt
	}
	err := volumes.Get(context.TODO(), client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22").ExtractInto(&s)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "304dc00909ac4d0da6c62d816bcb3459", s.TenantID)
	th.AssertEquals(t, "centos", s.Volume.VolumeImageMetadata["image_name"])

	err = volumes.Get(context.TODO(), client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22").ExtractInto(s)
	if err == nil {
		t.Errorf("Expected error when providing non-pointer struct")
	}
}

func TestCreateFromBackup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateVolumeFromBackupResponse(t)

	options := volumes.CreateOpts{
		Name:     "vol-001",
		BackupID: "20c792f0-bb03-434f-b653-06ef238e337e",
	}

	v, err := volumes.Create(context.TODO(), client.ServiceClient(), options).Extract()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, v.Size, 30)
	th.AssertEquals(t, v.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertEquals(t, *v.BackupID, "20c792f0-bb03-434f-b653-06ef238e337e")
}
