package msgutil

type Data map[string]interface{}

type Msg struct {
	Data Data
}

func NewMessage() Msg {
	return Msg{
		Data: make(Data),
	}
}

func (m Msg) Set(key string, value interface{}) Msg {
	m.Data[key] = value
	return m
}

func (m Msg) Done() Data {
	return m.Data
}

func RequestBodyParseError() Data {
	return NewMessage().Set("error", "Failed to parse request body").Done()
}

func SomethingWentWrong() Data {
	return NewMessage().Set("error", "Something went wrong").Done()
}

func InvalidID() Data {
	return NewMessage().Set("error", "Invalid ID").Done()
}

func RecordNotFound() Data {
	return NewMessage().Set("error", "Record not found").Done()
}
