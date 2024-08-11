/*

create table hashf as
ID varchar Primary Key,
data jsonb

HASH INDEX(ID)
*/

type Struct[T any] struct {
	ID *ID[T]
}

type ID[T any] string

func Structf[T any](format string, args ...any) Struct[T] {
	id := ID[T](fmt.Sprintf(format, args...))
	return Struct[T]{ID: &id}
}

func (h Struct[T]) Get(ptr *T, pairs ...any) error {
	var data T
	meta := reflect.ValueOf(&data)
	for i := 0; i < meta.NumField(); i++ {
		field := meta.Field(i)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.New(field.Type().Elem()))
		}
		fmt.Println(field.Kind())
	}

	return nil
}

type Encoder interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

func (h Struct[T]) Put(pairs ...any) error {
	if len(pairs)%2 != 0 {
		return fmt.Errorf("hash.Put: invalid number of arguments")
	}
	var data T
	meta := reflect.ValueOf(data)
	for i := 0; i < len(pairs); i += 2 {
		path := pairs[i]
		val := pairs[i+1]

		fmt.Println(meta.Kind())

		if key, ok := path.(string); ok {
			keys := strings.Split(key, ".")
			for i := 0; i < len(keys); i++ {
				fmt.Println("part", keys[i])
			}
			key = strings.Join(keys, ".")
			field := meta.FieldByName(key)
			if !field.IsValid() {
				return fmt.Errorf("hash.Put: invalid key %s", key)
			}
			if field.Kind() == reflect.Struct {
				fmt.Println("..skipping struct")
				continue
			}
			key := fmt.Sprint("'", key, "'")

			fmt.Println(field, val)

		} else {
			return fmt.Errorf("hash.Put: invalid key type, must be string")
		}

	}

	// var data T
	// meta := reflect.ValueOf(data).Elem()
	// for i := 0; i < meta.NumField(); i++ {
	// 	field := meta.Field(i)
	// 	// if field.Kind() == reflect.Ptr {
	// 	// 	field.Set(reflect.New(field.Type().Elem()))
	// 	// }

	// 	fmt.Print(field.Type().Name(), " ", field.Type(),
	// 		"\n")
	// }
	return nil
}

func (h Struct[T]) GetAll(ptr *T) error {
	return nil
}

// func (h Hash) Set()

// func (h Hash) GetString(id ID, path string) string { return "" }
// func (h Hash) GetInt64(id ID, path string) string  { return "" }
// func (h Hash) GetUInt64(id ID, path string) string { return "" }

// func CreateLinkedList(nams string) *LinkedList {
// 	list := &LinkedList{}

// 	// sql := `create table name as
// 	// ID Primary Key,
// 	// next_id int,
// 	// data jsonb

// 	return list
// }

type Player struct {
	Name      string
	LastMatch time.Time `json:"last_match"`
	Stats     Stats     `json:"stats"`
}

type Stats struct {
	Stamina    int64 `json:"stamina"`
	Nickness   int64 `json:"nickness"`
	LineShots  int64
	BoastShots int64
	DropShots  int64
	LobShots   int64
	CrossShots int64
	Power      int64
	Technique  int64
	Speed      int64
}

type PlayerWithShotsOnly struct {
	Stats StatsShotsOnly `json:"stats"`
}

type StatsShotsOnly struct {
	LineShots  int64 `json:"line_shots"`
	BoastShots int64 `json:"boast_shots"`
	DropShots  int64 `json:"drop_shots"`
	LobShots   int64 `json:"lob_shots"`
	CrossShots int64 `json:"cross_shots"`
}

func main() {
	// var p Player
	h := db.Structf[Player]("random:%d", 1)
	// h.Put(
	// ...
	err := h.Put(
		"LastMatch", time.Now().UTC(),
		"Stats.Stamina", 10,
		"Stats.Nickness", 10,
		"Stats.LineShots", 10,
	)
	fmt.Println(err)
	// h.Get(&p,
	// 	"stats.stamina",
	// 	"stats.nickness",
	// )
	var pso PlayerWithShotsOnly
	s := db.Structf[PlayerWithShotsOnly]("random:%d", 1)
	s.GetAll(&pso)
	// s.GetOne(&pso)
	// s.GetOne(&pso, "Stats.Stamina", "Stats.Nickness")
	// s.GetOneWithContext(context.TODO(), &pso)
	// s.GetMany(&pso[])
	// s.GetManyWithContext(context.TODO(), &pso[])

