package album

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/LMedez/go-api-rest/internal/entity"
	"github.com/LMedez/go-api-rest/pkg/log"
	"google.golang.org/api/iterator"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	// Get returns the album with the specified album ID.
	Get(ctx context.Context, id string) (entity.Album, error)
	// Count returns the number of albums.
	Count(ctx context.Context) (int, error)
	// Create saves a new album in the storage.
	Create(ctx context.Context, album entity.Album) error
	// Update updates the album with given ID in the storage.
	Update(ctx context.Context, album entity.Album) error
	// Delete removes the album with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists albums in database
type repository struct {
	firestore *firestore.Client
	logger    log.Logger
}

// NewRepository creates a new album repository
func NewRepository(firestoreClient *firestore.Client, logger log.Logger) Repository {
	return repository{firestoreClient, logger}
}

// Get reads the album with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Album, error) {
	var album entity.Album
	doc, err := r.firestore.Collection("albums").Doc(id).Get(ctx)
	if err != nil {
		return album, err
	}
	err = doc.DataTo(&album)
	if err != nil {
		return entity.Album{}, err
	}
	return album, err
}

// Create saves a new album record in the database.
// It returns the ID of the newly inserted album record.
func (r repository) Create(ctx context.Context, album entity.Album) error {
	_, _, err := r.firestore.Collection("albums").Add(ctx, album)
	return err
}

// Update saves the changes to an album in the database.
func (r repository) Update(ctx context.Context, album entity.Album) error {
	_, err := r.firestore.Collection("albums").Doc(album.ID).Set(ctx, album)
	return err
}

// Delete deletes an album with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	_, err := r.firestore.Collection("albums").Doc(id).Delete(ctx)
	return err
}

// Count returns the number of the album records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	var err error
	docs := r.firestore.Collection("albums").Documents(ctx)
	for {
		_, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		count++
	}
	return count, err
}
