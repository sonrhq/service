package service

import context "context"

// IdentityKeeper is the expected interface for the identity keeper.
type IdentityKeeper interface {
    // DeriveAccount derives an account for an identity.
    DeriveAccount(ctx context.Context, identityID string) (string, error)


    // LinkCredential links a credential to an identity.
    LinkCredential(ctx context.Context, identityID string) error

    // LinkPersona links a persona to an account and identity.
    LinkPersona(ctx context.Context, identityID string) error
}
