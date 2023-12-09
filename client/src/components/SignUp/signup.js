
import { Form, Button, Alert } from "react-bootstrap";
import axios from "axios";
import {useNavigate} from "react-router-dom";
import {useState} from "react";
import {toast, ToastContainer} from "react-toastify";

export const SignUp = () => {
    const notify = (text) => toast.error(text, {
        position: "top-right",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "colored",
    });
    const navigate = useNavigate();
    const [loading, isLoading] = useState(false)
    const [credentials, setCredentials] = useState({email: '', password: ''})
    const submitHandler = async e => {
        e.preventDefault()
        isLoading(true)
        await axios.post('https://url-shortener-kjie.onrender.com/api/auth/signup', {
            email: credentials.email,
            password: credentials.password
        }, {withCredentials: true})
            .then(function(response) {
                if (response.status === 200) {
                    localStorage.setItem("token", response.data['message'])
                    navigate('/home')
                }

            })
            .catch(function (err) {
                isLoading(false)
                notify(err.response.data.message)
            })
    }

    return (
        <>
            <ToastContainer
                position="top-right"
                autoClose={5000}
                hideProgressBar={false}
                newestOnTop={false}
                closeOnClick
                rtl={false}
                pauseOnFocusLoss
                draggable
                pauseOnHover
                theme="colored"
            />
            <div
                className="sign-in__wrapper"

                style={{ width: `450px`, height: `auto`, position: `absolute`, top: `50%`, left: `50%`, transform: `translate(-50%, -50%)`}}
            >
                {/* Overlay */}
                <div className="sign-in__backdrop"></div>
                {/* Form */}
                <Form className="shadow p-4 bg-white rounded" method="POST" onSubmit={submitHandler}>
                    {/* Header */}
                    <div className="h4 mb-2 text-center">Регистрация</div>
                    <Form.Group className="mb-2" controlId="email">
                        <Form.Label>Email</Form.Label>
                        <Form.Control
                            type="email"
                            placeholder="Email"
                            value={credentials.email}
                            onChange={e => setCredentials({ ...credentials, email: e.target.value })}
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-2" controlId="password">
                        <Form.Label>Password</Form.Label>
                        <Form.Control
                            type="password"
                            placeholder="Password"
                            value={credentials.password}
                            onChange={e => setCredentials({ ...credentials, password: e.target.value })}
                            required
                        />
                        <Form.Text id="password" muted>
                            Пароль должен содержать латинские буквы и цифры, а также должен состоять из 6 до 18 символов
                        </Form.Text>
                    </Form.Group>
                    {!loading ? (
                        <Button className="w-100" variant="primary" type="submit">
                            Регистрация
                        </Button>
                    ) : (
                        <Button className="w-100" variant="primary" type="submit">
                            Регистрируемся...
                        </Button>
                    )}


                    <div className="d-grid justify-content-end">
                        <a
                            href="/signin"
                            className="text-muted"
                            variant="link"
                            style = {{marginTop: `10px`}}
                        >
                            Вход
                        </a>
                    </div>

                </Form>

            </div>
        </>
    );
}