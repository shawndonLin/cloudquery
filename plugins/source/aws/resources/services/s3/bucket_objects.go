package s3

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func bucketObjects() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_objects",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html`,
		Resolver:    fetchS3BucketObjects,
		Transform:   transformers.TransformWithStruct(&types.Object{}, transformers.WithPrimaryKeyComponents("Key")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "bucket_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}
func fetchS3BucketObjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3
	region := parent.Get("region").(*scalar.String)
	if region == nil {
		return nil
	}
	objectsOutput, err := svc.ListObjectsV2(ctx, &s3.ListObjectsV2Input{Bucket: r.Name}, func(o *s3.Options) {
		o.Region = region.Value
	})
	if err != nil {
		if client.IsAWSError(err, "NoSuchBucketObjectConfiguration", "NoSuchBucket") {
			return nil
		}
		return err
	}
	if objectsOutput != nil {
		res <- objectsOutput.Contents
	}
	return nil
}
