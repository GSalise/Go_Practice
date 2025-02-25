package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"

	"github.com/GSalise/GQLPractice/graph/model"

	"github.com/google/uuid"
)

// UpsertCharacter is the resolver for the upsertCharacter field.
func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	// Sets the provided id if it is provided
	id := input.ID

	//Create a new Character object
	var character model.Character
	character.Name = input.Name
	character.CliqueType = input.CliqueType

	// n := len(r.Resolver.CharacterStore)
	// if n == 0 {
	// 	r.Resolver.CharacterStore = make(map[string]model.Character)
	// }

	//If an id is provided, update the character
	if input.ID != nil {
		// Fetch the existing character from the database
		err := r.DB.QueryRow(ctx, `
			SELECT id, name, is_hero, clique_type
			FROM characters
			WHERE id = $1
		`, id).Scan(&character.ID, &character.Name, &character.IsHero, &character.CliqueType)
		if err != nil {
			return nil, fmt.Errorf("character not found: %v", err)
		}

		// Update the character's fields
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		}
		character.Name = input.Name
		character.CliqueType = input.CliqueType

		// Update the character in the database
		_, err = r.DB.Exec(ctx, `
			UPDATE characters
			SET name = $1, is_hero = $2, clique_type = $3
			WHERE id = $4
		`, character.Name, character.IsHero, character.CliqueType, id)
		if err != nil {
			return nil, fmt.Errorf("failed to update character: %v", err)
		}
	} else {
		// If no ID is provided, create a new character
		newID := uuid.New().String()
		character.ID = newID
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		}

		// Insert the new character into the database
		_, err := r.DB.Exec(ctx, `
			INSERT INTO characters (id, name, is_hero, clique_type)
			VALUES ($1, $2, $3, $4)
		`, character.ID, character.Name, character.IsHero, character.CliqueType)
		if err != nil {
			return nil, fmt.Errorf("failed to insert character: %v", err)
		}
	}

	// Return the updated or new character
	return &character, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	var character model.Character

	// Fetch the character from the database
	err := r.DB.QueryRow(ctx, `
		SELECT id, name, is_hero, clique_type
		FROM characters
		WHERE id = $1
	`, id).Scan(&character.ID, &character.Name, &character.IsHero, &character.CliqueType)
	if err != nil {
		return nil, fmt.Errorf("character not found: %v", err)
	}

	return &character, nil
}

// Pogues is the resolver for the pogues field.
func (r *queryResolver) Pogues(ctx context.Context) ([]*model.Character, error) {
	var pogues []*model.Character

	rows, err := r.DB.Query(ctx, `
		SELECT id, name, is_hero, clique_type
		FROM characters
		WHERE clique_type = $1
	`, model.CliqueTypePogues)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch pogues: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var character model.Character
		err := rows.Scan(&character.ID, &character.Name, &character.IsHero, &character.CliqueType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan character: %v", err)
		}

		pogues = append(pogues, &character)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error inserting rows: %v", err)
	}

	return pogues, nil

}

// Kooks is the resolver for the kooks field.
func (r *queryResolver) Kooks(ctx context.Context) ([]*model.Character, error) {
	var kooks []*model.Character

	rows, err := r.DB.Query(ctx, `
		SELECT id, name, is_hero, clique_type
		FROM characters
		WHERE clique_type = $1
	`, model.CliqueTypeKooks)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch kooks: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var character model.Character
		err := rows.Scan(&character.ID, &character.Name, &character.IsHero, &character.CliqueType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan character: %v", err)
		}

		kooks = append(kooks, &character)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error inserting rows: %v", err)
	}

	return kooks, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
