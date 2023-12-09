import Navbar from 'react-bootstrap/Navbar';
import {Container, Nav} from "react-bootstrap";
export const NavBar = () => {
    return (
        <Navbar bg="dark" data-bs-theme="dark" style={{ marginBottom: `15px` }}>
            <Container>
                <Navbar.Brand href="/home">Url-Shortener</Navbar.Brand>
                <Nav className="me-auto">
                    <Nav.Link href="/home">Home</Nav.Link>
                    <Nav.Link href="/about">About</Nav.Link>
                    <Nav.Link href="/signin">Logout</Nav.Link>
                </Nav>
            </Container>
        </Navbar>
    );
}