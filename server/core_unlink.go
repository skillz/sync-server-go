// //
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"

	"go.uber.org/zap"
)

func UnlinkCustom(ctx context.Context, logger *zap.Logger, db *sql.DB, id uuid.UUID, customID string) error {
//	if customID == "" {
//		return status.Error(codes.InvalidArgument, "An ID must be supplied.")
//	}
//
//	res, err := db.ExecContext(ctx, `UPDATE users SET custom_id = NULL, update_time = now()
//WHERE id = $1
//AND custom_id = $2
//AND ((facebook_id IS NOT NULL
//      OR facebook_instant_game_id IS NOT NULL
//      OR google_id IS NOT NULL
//      OR gamecenter_id IS NOT NULL
//      OR steam_id IS NOT NULL
//      OR email IS NOT NULL)
//     OR
//     EXISTS (SELECT id FROM user_device WHERE user_id = $1 LIMIT 1))`, id, customID)
//
//	if err != nil {
//		logger.Error("Could not unlink custom ID.", zap.Error(err), zap.Any("input", customID))
//		return status.Error(codes.Internal, "Error while trying to unlink custom ID.")
//	} else if count, _ := res.RowsAffected(); count == 0 {
//		return status.Error(codes.PermissionDenied, "Cannot unlink last account identifier. Check profile exists and is not last link.")
//	}
	return nil
}

func UnlinkDevice(ctx context.Context, logger *zap.Logger, db *sql.DB, id uuid.UUID, deviceID string) error {
//	if deviceID == "" {
//		return status.Error(codes.InvalidArgument, "A device ID must be supplied.")
//	}
//
//	tx, err := db.BeginTx(ctx, nil)
//	if err != nil {
//		logger.Error("Could not begin database transaction.", zap.Error(err))
//		return status.Error(codes.Internal, "Could not unlink Device ID.")
//	}
//
//	err = ExecuteInTx(ctx, tx, func() error {
//		res, err := tx.ExecContext(ctx, `DELETE FROM user_device WHERE id = $2 AND user_id = $1
//AND (EXISTS (SELECT id FROM users WHERE id = $1 AND
//    (facebook_id IS NOT NULL
//     OR facebook_instant_game_id IS NOT NULL
//     OR google_id IS NOT NULL
//     OR gamecenter_id IS NOT NULL
//     OR steam_id IS NOT NULL
//     OR email IS NOT NULL
//     OR custom_id IS NOT NULL))
//   OR EXISTS (SELECT id FROM user_device WHERE user_id = $1 AND id <> $2 LIMIT 1))`, id, deviceID)
//		if err != nil {
//			logger.Debug("Could not unlink device ID.", zap.Error(err), zap.Any("input", deviceID))
//			return err
//		}
//		if count, _ := res.RowsAffected(); count == 0 {
//			return StatusError(codes.PermissionDenied, "Cannot unlink last account identifier. Check profile exists and is not last link.", ErrRowsAffectedCount)
//		}
//
//		res, err = tx.ExecContext(ctx, "UPDATE users SET update_time = now() WHERE id = $1", id)
//		if err != nil {
//			logger.Debug("Could not unlink device ID.", zap.Error(err), zap.Any("input", deviceID))
//			return err
//		}
//		if count, _ := res.RowsAffected(); count == 0 {
//			return StatusError(codes.PermissionDenied, "Cannot unlink last account identifier. Check profile exists and is not last link.", ErrRowsAffectedCount)
//		}
//
//		return nil
//	})
//
//	if err != nil {
//		if e, ok := err.(*statusError); ok {
//			return e.Status()
//		}
//		logger.Error("Error in database transaction.", zap.Error(err))
//		return status.Error(codes.Internal, "Could not unlink device ID.")
//	}
	return nil
}

func UnlinkEmail(ctx context.Context, logger *zap.Logger, db *sql.DB, id uuid.UUID, email string) error {
//	if email == "" {
//		return status.Error(codes.InvalidArgument, "Both email and password must be supplied.")
//	}
//	cleanEmail := strings.ToLower(email)
//
//	res, err := db.ExecContext(ctx, `UPDATE users SET email = NULL, password = NULL, update_time = now()
//WHERE id = $1
//AND email = $2
//AND ((facebook_id IS NOT NULL
//      OR facebook_instant_game_id IS NOT NULL
//      OR google_id IS NOT NULL
//      OR gamecenter_id IS NOT NULL
//      OR steam_id IS NOT NULL
//      OR custom_id IS NOT NULL)
//     OR
//     EXISTS (SELECT id FROM user_device WHERE user_id = $1 LIMIT 1))`, id, cleanEmail)
//
//	if err != nil {
//		logger.Error("Could not unlink email.", zap.Error(err), zap.Any("input", email))
//		return status.Error(codes.Internal, "Error while trying to unlink email.")
//	} else if count, _ := res.RowsAffected(); count == 0 {
//		return status.Error(codes.PermissionDenied, "Cannot unlink last account identifier. Check profile exists and is not last link.")
//	}
	return nil
}
