// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	Group
	Entity
	Alias
	Persona
	EntityStorageEntry
	PersonaIndexEntry
*/
package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Group represents an identity group.
type Group struct {
	// ID is the unique identifier for this group
	ID string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Name is the unique name for this group
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Policies are the vault policies to be granted to members of this group
	Policies []string `protobuf:"bytes,3,rep,name=policies" json:"policies,omitempty"`
	// ParentGroupIDs are the identifiers of those groups to which this group is a
	// member of. These will serve as references to the parent group in the
	// hierarchy.
	ParentGroupIDs []string `protobuf:"bytes,4,rep,name=parent_group_ids,json=parentGroupIds" json:"parent_group_ids,omitempty"`
	// MemberEntityIDs are the identifiers of entities which are members of this
	// group
	MemberEntityIDs []string `protobuf:"bytes,5,rep,name=member_entity_ids,json=memberEntityIDs" json:"member_entity_ids,omitempty"`
	// Metadata represents the custom data tied with this group
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// CreationTime is the time at which this group was created
	CreationTime *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// LastUpdateTime is the time at which this group was last modified
	LastUpdateTime *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	// ModifyIndex tracks the number of updates to the group. It is useful to detect
	// updates to the groups.
	ModifyIndex uint64 `protobuf:"varint,9,opt,name=modify_index,json=modifyIndex" json:"modify_index,omitempty"`
	// BucketKeyHash is the MD5 hash of the storage bucket key into which this
	// group is stored in the underlying storage. This is useful to find all
	// the groups belonging to a particular bucket during invalidation of the
	// storage key.
	BucketKeyHash string `protobuf:"bytes,10,opt,name=bucket_key_hash,json=bucketKeyHash" json:"bucket_key_hash,omitempty"`
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Group) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetPolicies() []string {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *Group) GetParentGroupIDs() []string {
	if m != nil {
		return m.ParentGroupIDs
	}
	return nil
}

func (m *Group) GetMemberEntityIDs() []string {
	if m != nil {
		return m.MemberEntityIDs
	}
	return nil
}

func (m *Group) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Group) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Group) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Group) GetModifyIndex() uint64 {
	if m != nil {
		return m.ModifyIndex
	}
	return 0
}

func (m *Group) GetBucketKeyHash() string {
	if m != nil {
		return m.BucketKeyHash
	}
	return ""
}

// Entity represents an entity that gets persisted and indexed.
// Entity is fundamentally composed of zero or many aliases.
type Entity struct {
	// Deprecated in favor of Aliases.
	// Personas are the identities that this entity is made of. This can be
	// empty as well to favor being able to create the entity first and then
	// incrementally adding personas.
	Personas []*Persona `protobuf:"bytes,1,rep,name=personas" json:"personas,omitempty"`
	// ID is the unique identifier of the entity which always be a UUID. This
	// should never be allowed to be updated.
	ID string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// Name is a unique identifier of the entity which is intended to be
	// human-friendly. The default name might not be human friendly since it
	// gets suffixed by a UUID, but it can optionally be updated, unlike the ID
	// field.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Metadata represents the explicit metadata which is set by the
	// clients.  This is useful to tie any information pertaining to the
	// aliases. This is a non-unique field of entity, meaning multiple
	// entities can have the same metadata set. Entities will be indexed based
	// on this explicit metadata. This enables virtual groupings of entities
	// based on its metadata.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// CreationTime is the time at which this entity is first created.
	CreationTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// LastUpdateTime is the most recent time at which the properties of this
	// entity got modified. This is helpful in filtering out entities based on
	// its age and to take action on them, if desired.
	LastUpdateTime *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	// MergedEntityIDs are the entities which got merged to this one. Entities
	// will be indexed based on all the entities that got merged into it. This
	// helps to apply the actions on this entity on the tokens that are merged
	// to the merged entities. Merged entities will be deleted entirely and
	// this is the only trackable trail of its earlier presence.
	MergedEntityIDs []string `protobuf:"bytes,7,rep,name=merged_entity_ids,json=mergedEntityIDs" json:"merged_entity_ids,omitempty"`
	// Policies the entity is entitled to
	Policies []string `protobuf:"bytes,8,rep,name=policies" json:"policies,omitempty"`
	// BucketKeyHash is the MD5 hash of the storage bucket key into which this
	// entity is stored in the underlying storage. This is useful to find all
	// the entities belonging to a particular bucket during invalidation of the
	// storage key.
	BucketKeyHash string   `protobuf:"bytes,9,opt,name=bucket_key_hash,json=bucketKeyHash" json:"bucket_key_hash,omitempty"`
	Aliases       []*Alias `protobuf:"bytes,11,rep,name=aliases" json:"aliases,omitempty"`
}

func (m *Entity) Reset()                    { *m = Entity{} }
func (m *Entity) String() string            { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()               {}
func (*Entity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Entity) GetPersonas() []*Persona {
	if m != nil {
		return m.Personas
	}
	return nil
}

func (m *Entity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Entity) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Entity) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Entity) GetMergedEntityIDs() []string {
	if m != nil {
		return m.MergedEntityIDs
	}
	return nil
}

func (m *Entity) GetPolicies() []string {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *Entity) GetBucketKeyHash() string {
	if m != nil {
		return m.BucketKeyHash
	}
	return ""
}

func (m *Entity) GetAliases() []*Alias {
	if m != nil {
		return m.Aliases
	}
	return nil
}

// Alias represents the alias that gets stored inside of the
// entity object in storage and also represents in an in-memory index of an
// alias object.
type Alias struct {
	// ID is the unique identifier that represents this alias
	ID string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// EntityID is the entity identifier to which this alias belongs to
	EntityID string `protobuf:"bytes,2,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	// MountType is the backend mount's type to which this alias belongs to.
	// This enables categorically querying aliases of specific backend types.
	MountType string `protobuf:"bytes,3,opt,name=mount_type,json=mountType" json:"mount_type,omitempty"`
	// MountAccessor is the backend mount's accessor to which this alias
	// belongs to.
	MountAccessor string `protobuf:"bytes,4,opt,name=mount_accessor,json=mountAccessor" json:"mount_accessor,omitempty"`
	// MountPath is the backend mount's path to which the Maccessor belongs to. This
	// field is not used for any operational purposes. This is only returned when
	// alias is read, only as a nicety.
	MountPath string `protobuf:"bytes,5,opt,name=mount_path,json=mountPath" json:"mount_path,omitempty"`
	// Metadata is the explicit metadata that clients set against an entity
	// which enables virtual grouping of aliases. Aliases will be indexed
	// against their metadata.
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Name is the identifier of this alias in its authentication source.
	// This does not uniquely identify a alias in Vault. This in conjunction
	// with MountAccessor form to be the factors that represent a alias in a
	// unique way. Aliases will be indexed based on this combined uniqueness
	// factor.
	Name string `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	// CreationTime is the time at which this alias was first created
	CreationTime *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// LastUpdateTime is the most recent time at which the properties of this
	// alias got modified. This is helpful in filtering out aliases based
	// on its age and to take action on them, if desired.
	LastUpdateTime *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	// MergedFromEntityIDs is the FIFO history of merging activity by entity IDs from
	// which this alias is transfered over to the entity to which it
	// currently belongs to.
	MergedFromEntityIDs []string `protobuf:"bytes,10,rep,name=merged_from_entity_ids,json=mergedFromEntityIDs" json:"merged_from_entity_ids,omitempty"`
}

func (m *Alias) Reset()                    { *m = Alias{} }
func (m *Alias) String() string            { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()               {}
func (*Alias) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Alias) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Alias) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

func (m *Alias) GetMountType() string {
	if m != nil {
		return m.MountType
	}
	return ""
}

func (m *Alias) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *Alias) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *Alias) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Alias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Alias) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Alias) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Alias) GetMergedFromEntityIDs() []string {
	if m != nil {
		return m.MergedFromEntityIDs
	}
	return nil
}

type Persona struct {
	ID                  string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	EntityID            string                     `protobuf:"bytes,2,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	MountType           string                     `protobuf:"bytes,3,opt,name=mount_type,json=mountType" json:"mount_type,omitempty"`
	MountAccessor       string                     `protobuf:"bytes,4,opt,name=mount_accessor,json=mountAccessor" json:"mount_accessor,omitempty"`
	MountPath           string                     `protobuf:"bytes,5,opt,name=mount_path,json=mountPath" json:"mount_path,omitempty"`
	Metadata            map[string]string          `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Name                string                     `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	CreationTime        *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	LastUpdateTime      *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	MergedFromEntityIDs []string                   `protobuf:"bytes,10,rep,name=merged_from_entity_ids,json=mergedFromEntityIDs" json:"merged_from_entity_ids,omitempty"`
}

func (m *Persona) Reset()                    { *m = Persona{} }
func (m *Persona) String() string            { return proto.CompactTextString(m) }
func (*Persona) ProtoMessage()               {}
func (*Persona) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Persona) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Persona) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

func (m *Persona) GetMountType() string {
	if m != nil {
		return m.MountType
	}
	return ""
}

func (m *Persona) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *Persona) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *Persona) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Persona) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Persona) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Persona) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Persona) GetMergedFromEntityIDs() []string {
	if m != nil {
		return m.MergedFromEntityIDs
	}
	return nil
}

// Deprecated. Retained for backwards compatibility.
type EntityStorageEntry struct {
	Personas        []*PersonaIndexEntry       `protobuf:"bytes,1,rep,name=personas" json:"personas,omitempty"`
	ID              string                     `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Name            string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Metadata        map[string]string          `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CreationTime    *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	LastUpdateTime  *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	MergedEntityIDs []string                   `protobuf:"bytes,7,rep,name=merged_entity_ids,json=mergedEntityIDs" json:"merged_entity_ids,omitempty"`
	Policies        []string                   `protobuf:"bytes,8,rep,name=policies" json:"policies,omitempty"`
	BucketKeyHash   string                     `protobuf:"bytes,9,opt,name=bucket_key_hash,json=bucketKeyHash" json:"bucket_key_hash,omitempty"`
}

func (m *EntityStorageEntry) Reset()                    { *m = EntityStorageEntry{} }
func (m *EntityStorageEntry) String() string            { return proto.CompactTextString(m) }
func (*EntityStorageEntry) ProtoMessage()               {}
func (*EntityStorageEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EntityStorageEntry) GetPersonas() []*PersonaIndexEntry {
	if m != nil {
		return m.Personas
	}
	return nil
}

func (m *EntityStorageEntry) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *EntityStorageEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntityStorageEntry) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EntityStorageEntry) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *EntityStorageEntry) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *EntityStorageEntry) GetMergedEntityIDs() []string {
	if m != nil {
		return m.MergedEntityIDs
	}
	return nil
}

func (m *EntityStorageEntry) GetPolicies() []string {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *EntityStorageEntry) GetBucketKeyHash() string {
	if m != nil {
		return m.BucketKeyHash
	}
	return ""
}

// Deprecated. Retained for backwards compatibility.
type PersonaIndexEntry struct {
	ID                  string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	EntityID            string                     `protobuf:"bytes,2,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	MountType           string                     `protobuf:"bytes,3,opt,name=mount_type,json=mountType" json:"mount_type,omitempty"`
	MountAccessor       string                     `protobuf:"bytes,4,opt,name=mount_accessor,json=mountAccessor" json:"mount_accessor,omitempty"`
	MountPath           string                     `protobuf:"bytes,5,opt,name=mount_path,json=mountPath" json:"mount_path,omitempty"`
	Metadata            map[string]string          `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Name                string                     `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	CreationTime        *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	LastUpdateTime      *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	MergedFromEntityIDs []string                   `protobuf:"bytes,10,rep,name=merged_from_entity_ids,json=mergedFromEntityIDs" json:"merged_from_entity_ids,omitempty"`
}

func (m *PersonaIndexEntry) Reset()                    { *m = PersonaIndexEntry{} }
func (m *PersonaIndexEntry) String() string            { return proto.CompactTextString(m) }
func (*PersonaIndexEntry) ProtoMessage()               {}
func (*PersonaIndexEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PersonaIndexEntry) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PersonaIndexEntry) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

func (m *PersonaIndexEntry) GetMountType() string {
	if m != nil {
		return m.MountType
	}
	return ""
}

func (m *PersonaIndexEntry) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *PersonaIndexEntry) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *PersonaIndexEntry) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PersonaIndexEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonaIndexEntry) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *PersonaIndexEntry) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *PersonaIndexEntry) GetMergedFromEntityIDs() []string {
	if m != nil {
		return m.MergedFromEntityIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*Group)(nil), "identity.Group")
	proto.RegisterType((*Entity)(nil), "identity.Entity")
	proto.RegisterType((*Alias)(nil), "identity.Alias")
	proto.RegisterType((*Persona)(nil), "identity.Persona")
	proto.RegisterType((*EntityStorageEntry)(nil), "identity.EntityStorageEntry")
	proto.RegisterType((*PersonaIndexEntry)(nil), "identity.PersonaIndexEntry")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x55, 0xcf, 0x4f, 0xdb, 0x4c,
	0x10, 0x55, 0xe2, 0xfc, 0xb0, 0x27, 0x10, 0x60, 0xbf, 0x4f, 0x95, 0x15, 0x44, 0x49, 0x91, 0x5a,
	0x05, 0xa4, 0x1a, 0x09, 0x0e, 0x6d, 0xe1, 0x50, 0x21, 0x15, 0x5a, 0x54, 0x55, 0x42, 0x2e, 0x3d,
	0x5b, 0x9b, 0x78, 0x49, 0x2c, 0x62, 0xaf, 0xb5, 0xbb, 0xae, 0xea, 0x3f, 0xb7, 0x87, 0xaa, 0xa7,
	0x5e, 0x2b, 0xf5, 0x56, 0x79, 0xd6, 0x4e, 0x5c, 0x4c, 0x29, 0x08, 0x0e, 0x1c, 0xb8, 0xd9, 0x33,
	0xb3, 0xcf, 0x3b, 0xef, 0xbd, 0x19, 0x43, 0x47, 0xa5, 0x31, 0x93, 0x4e, 0x2c, 0xb8, 0xe2, 0xc4,
	0x0c, 0x7c, 0x16, 0xa9, 0x40, 0xa5, 0xbd, 0xf5, 0x31, 0xe7, 0xe3, 0x29, 0xdb, 0xc6, 0xf8, 0x30,
	0x39, 0xdb, 0x56, 0x41, 0xc8, 0xa4, 0xa2, 0x61, 0xac, 0x4b, 0x37, 0xbe, 0x1b, 0xd0, 0x7c, 0x2b,
	0x78, 0x12, 0x93, 0x2e, 0xd4, 0x03, 0xdf, 0xae, 0xf5, 0x6b, 0x03, 0xcb, 0xad, 0x07, 0x3e, 0x21,
	0xd0, 0x88, 0x68, 0xc8, 0xec, 0x3a, 0x46, 0xf0, 0x99, 0xf4, 0xc0, 0x8c, 0xf9, 0x34, 0x18, 0x05,
	0x4c, 0xda, 0x46, 0xdf, 0x18, 0x58, 0xee, 0xec, 0x9d, 0x0c, 0x60, 0x39, 0xa6, 0x82, 0x45, 0xca,
	0x1b, 0x67, 0x78, 0x5e, 0xe0, 0x4b, 0xbb, 0x81, 0x35, 0x5d, 0x1d, 0xc7, 0xcf, 0x1c, 0xfb, 0x92,
	0x6c, 0xc1, 0x4a, 0xc8, 0xc2, 0x21, 0x13, 0x9e, 0xbe, 0x25, 0x96, 0x36, 0xb1, 0x74, 0x49, 0x27,
	0x0e, 0x31, 0x9e, 0xd5, 0xbe, 0x02, 0x33, 0x64, 0x8a, 0xfa, 0x54, 0x51, 0xbb, 0xd5, 0x37, 0x06,
	0x9d, 0x9d, 0x35, 0xa7, 0xe8, 0xce, 0x41, 0x44, 0xe7, 0x43, 0x9e, 0x3f, 0x8c, 0x94, 0x48, 0xdd,
	0x59, 0x39, 0x79, 0x0d, 0x8b, 0x23, 0xc1, 0xa8, 0x0a, 0x78, 0xe4, 0x65, 0x6d, 0xdb, 0xed, 0x7e,
	0x6d, 0xd0, 0xd9, 0xe9, 0x39, 0x9a, 0x13, 0xa7, 0xe0, 0xc4, 0x39, 0x2d, 0x38, 0x71, 0x17, 0x8a,
	0x03, 0x59, 0x88, 0xbc, 0x81, 0xe5, 0x29, 0x95, 0xca, 0x4b, 0x62, 0x9f, 0x2a, 0xa6, 0x31, 0xcc,
	0x7f, 0x62, 0x74, 0xb3, 0x33, 0x9f, 0xf0, 0x08, 0xa2, 0x3c, 0x81, 0x85, 0x90, 0xfb, 0xc1, 0x59,
	0xea, 0x05, 0x91, 0xcf, 0xbe, 0xd8, 0x56, 0xbf, 0x36, 0x68, 0xb8, 0x1d, 0x1d, 0x3b, 0xce, 0x42,
	0xe4, 0x19, 0x2c, 0x0d, 0x93, 0xd1, 0x39, 0x53, 0xde, 0x39, 0x4b, 0xbd, 0x09, 0x95, 0x13, 0x1b,
	0x90, 0xf5, 0x45, 0x1d, 0x7e, 0xcf, 0xd2, 0x77, 0x54, 0x4e, 0x7a, 0xfb, 0xb0, 0xf8, 0x47, 0xb3,
	0x64, 0x19, 0x8c, 0x73, 0x96, 0xe6, 0xa2, 0x65, 0x8f, 0xe4, 0x7f, 0x68, 0x7e, 0xa6, 0xd3, 0xa4,
	0x90, 0x4d, 0xbf, 0xec, 0xd5, 0x5f, 0xd6, 0x36, 0x7e, 0x1a, 0xd0, 0xd2, 0xbc, 0x92, 0xe7, 0x60,
	0xc6, 0x4c, 0x48, 0x1e, 0x51, 0x69, 0xd7, 0x90, 0xd4, 0x95, 0x39, 0xa9, 0x27, 0x3a, 0xe3, 0xce,
	0x4a, 0x72, 0x67, 0xd4, 0x2b, 0xce, 0x30, 0x4a, 0xce, 0xd8, 0x2b, 0xe9, 0xd4, 0x40, 0xc8, 0xc7,
	0x73, 0x48, 0xfd, 0xd9, 0xeb, 0x0b, 0xd5, 0xbc, 0x03, 0xa1, 0x5a, 0x37, 0x16, 0x0a, 0x6d, 0x29,
	0xc6, 0xcc, 0x2f, 0xdb, 0xb2, 0x5d, 0xd8, 0x32, 0x4b, 0xcc, 0x6d, 0x59, 0x1e, 0x04, 0xf3, 0xc2,
	0x20, 0x5c, 0xa2, 0xa6, 0x75, 0x89, 0x9a, 0x64, 0x13, 0xda, 0x74, 0x1a, 0x50, 0xc9, 0xa4, 0xdd,
	0x41, 0xc6, 0x96, 0xe6, 0x8c, 0x1d, 0x64, 0x09, 0xb7, 0xc8, 0xdf, 0x4e, 0xf8, 0xaf, 0x06, 0x34,
	0x11, 0xaf, 0x32, 0xe2, 0xab, 0x60, 0xcd, 0x5a, 0xcd, 0xcf, 0x99, 0x2c, 0xef, 0x91, 0xac, 0x01,
	0x84, 0x3c, 0x89, 0x94, 0x97, 0x6d, 0x96, 0x5c, 0x6b, 0x0b, 0x23, 0xa7, 0x69, 0xcc, 0xc8, 0x53,
	0xe8, 0xea, 0x34, 0x1d, 0x8d, 0x98, 0x94, 0x5c, 0xd8, 0x0d, 0xdd, 0x24, 0x46, 0x0f, 0xf2, 0xe0,
	0x1c, 0x25, 0xa6, 0x6a, 0x82, 0xc2, 0x16, 0x28, 0x27, 0x54, 0x4d, 0xae, 0x1e, 0x6f, 0xbc, 0xf4,
	0x5f, 0x5d, 0x53, 0xb8, 0xb0, 0x5d, 0x72, 0x61, 0xc5, 0x49, 0xe6, 0x1d, 0x38, 0xc9, 0xba, 0xb1,
	0x93, 0x76, 0xe1, 0x51, 0xee, 0xa4, 0x33, 0xc1, 0xc3, 0xb2, 0x9d, 0x00, 0xbd, 0xf2, 0x9f, 0xce,
	0x1e, 0x09, 0x1e, 0xce, 0x2c, 0x75, 0x3b, 0x8d, 0xbf, 0x19, 0xd0, 0xce, 0x07, 0xf7, 0x1e, 0xaa,
	0xbc, 0x5f, 0x51, 0x79, 0xbd, 0xb2, 0x6f, 0x1e, 0x74, 0xbe, 0x8e, 0xce, 0x3f, 0x0c, 0x20, 0x1a,
	0xea, 0xa3, 0xe2, 0x82, 0x8e, 0x99, 0x86, 0x78, 0x51, 0x59, 0xe8, 0xab, 0x15, 0x82, 0xf1, 0x57,
	0x93, 0x93, 0x7b, 0xa3, 0xd5, 0x7e, 0x54, 0x59, 0xed, 0x5b, 0x17, 0x57, 0x7b, 0xf9, 0x32, 0x0f,
	0x6b, 0xfe, 0x0e, 0x7f, 0xda, 0xbf, 0x0c, 0x58, 0xa9, 0xe8, 0x77, 0x0f, 0x27, 0xfc, 0xb0, 0x32,
	0xe1, 0x9b, 0x57, 0x18, 0xf0, 0x61, 0xd6, 0xaf, 0xa1, 0xfd, 0xb0, 0x85, 0xb7, 0xda, 0xfd, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0xf3, 0x0c, 0x03, 0x2e, 0xdb, 0x0b, 0x00, 0x00,
}
