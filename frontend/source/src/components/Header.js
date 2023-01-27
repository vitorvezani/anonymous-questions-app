import { Link } from "react-router-dom";

import styles from "./Header.module.css";

const Header = () => {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Anonymous Question App</h1>
      <Link to="/" className={styles.link}>
        View Questions
      </Link>
      <Link to="/add-question" className={styles.link}>
        Add Question
      </Link>
    </div>
  );
};

export default Header;
