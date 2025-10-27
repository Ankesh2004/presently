package api

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
	Role     string             `bson:"role" json:"role"` // student or teacher
}

type Classroom struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name         string               `bson:"name" json:"name"`
	InstructorID primitive.ObjectID   `bson:"instructorId" json:"instructorId"`
	StudentIDs   []primitive.ObjectID `bson:"studentIds" json:"studentIds"`
	UniqueCode   string               `bson:"uniqueCode" json:"uniqueCode"`
}

type AttendanceSession struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ClassroomID primitive.ObjectID `bson:"classroomId" json:"classroomId"`
	StartTime   time.Time          `bson:"startTime" json:"startTime"`
	EndTime     time.Time          `bson:"endTime" json:"endTime"`
	QRCodeData  string             `bson:"qrCodeData" json:"qrCodeData"` //session ID
}

type AttendanceRecord struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SessionID primitive.ObjectID `bson:"sessionId" json:"sessionId"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	Status    string             `bson:"status" json:"status"` // PRESENT or ABSENT
}
