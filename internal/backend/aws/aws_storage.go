package aws

import "QuickBin/internal/backend"

type AWSStorage struct {
	// AWS S3 Bucket
	bucket string
	// AWS S3 Region
	region string
	// AWS S3 Access Key
	accessKey string
	// AWS S3 Secret Key
	secretKey string
}

func NewAWSStorage(options map[string]interface{}) backend.StorageImplement {
	bucket := options["bucket"].(string)
	if bucket == "" {
		panic("bucket is required for aws storage")
	}
	region := options["region"].(string)
	if region == "" {
		panic("region is required for aws storage")
	}
	accessKey := options["access-key"].(string)
	if accessKey == "" {
		panic("access-key is required for aws storage")
	}
	secretKey := options["secret-key"].(string)
	if secretKey == "" {
		panic("secret-key is required for aws storage")
	}
	return &AWSStorage{
		bucket:    bucket,
		region:    region,
		accessKey: accessKey,
		secretKey: secretKey,
	}
}

func (as *AWSStorage) Save(key string, data []byte) error {
	// Save the data to the AWS S3 Bucket
	return nil
}

func (as *AWSStorage) Get(key string) ([]byte, error) {
	// Get the data from the AWS S3 Bucket
	return nil, nil
}

func (as *AWSStorage) Delete(key string) error {
	// Delete the data from the AWS S3 Bucket
	return nil
}
